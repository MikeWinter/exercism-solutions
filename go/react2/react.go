package react

type manager struct {
	deps map[Cell][]*compute
}
type input struct {
	val int
	m   *manager
}
type compute struct {
	val       int
	m         *manager
	next      int
	update    func(int, int) int
	deps      []Cell
	updated   bool
	callbacks map[int]func(int)
}
type cancel struct {
	cell  *compute
	index int
}

// New - create a new Reactor to manage Cells
func New() Reactor {
	return &manager{deps: make(map[Cell][]*compute)}
}

func (m *manager) CreateInput(val int) InputCell {
	i := &input{val: val, m: m}
	m.deps[i] = []*compute{} // Add new InputCell to dependency list with no dependecies
	return i
}

func (m *manager) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	res := f(c.Value())
	g := func(x, _ int) int { return f(x) }     // Wrap function
	cell := &compute{val: res, m: m, update: g, // Make new ComputeCell and add passed in Cell as a dependency
		deps: []Cell{c}, callbacks: make(map[int]func(int))}
	m.deps[cell] = []*compute{}
	m.deps[c] = append(m.deps[c], cell) // Add new ComputeCell as a dependency of the passed in Cell
	return cell
}

func (m *manager) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	res := f(c1.Value(), c2.Value())
	cell := &compute{val: res, m: m, update: f, // Make new ComputeCell and add passed in Cells as dependenies
		deps: []Cell{c1, c2}, callbacks: make(map[int]func(int))}
	m.deps[cell] = []*compute{}
	m.deps[c1] = append(m.deps[c1], cell) // Add new ComputeCell as a dependency of the first passed in Cell
	m.deps[c2] = append(m.deps[c2], cell) // Add new ComputeCell as a dependency of the second passed in Cell
	return cell
}

func (cell *input) Value() int {
	return cell.val
}

func (cell *input) SetValue(val int) {
	if cell.val == val {
		return
	}
	cell.val = val
	prev := cell.m.snapshot()
	cell.m.updateDeps(cell)
	cell.m.markChanged(prev)
	cell.m.callCallbacks(cell)
}

func (cell *compute) Value() int {
	return cell.val
}

func (cell *compute) AddCallback(f func(int)) Canceler {
	defer func() { cell.next++ }() // increment 'next' counter after function returns (bad security to do simple inc)
	cell.callbacks[cell.next] = f  // add callback at current 'next' position
	return cancel{cell: cell, index: cell.next}
}

func (remove cancel) Cancel() {
	delete(remove.cell.callbacks, remove.index)
}

func (m *manager) updateDeps(c Cell) {
	// update dependencies recursively
	for _, cell := range m.deps[c] {
		dep1, dep2 := cell.deps[0].Value(), 0
		if len(cell.deps) == 2 {
			dep2 = cell.deps[1].Value()
		}
		old := cell.val
		cell.val = cell.update(dep1, dep2)
		if cell.val != old {
			m.updateDeps(cell)
		}
	}
}

func (m *manager) snapshot() map[Cell]int {
	state := make(map[Cell]int)
	for cell := range m.deps {
		state[cell] = cell.Value()
	}
	return state
}

func (m *manager) markChanged(prev map[Cell]int) {
	for cell := range m.deps {
		if c, ok := cell.(*compute); ok && c.val != prev[cell] {
			c.updated = true
		}
	}
}

func (m *manager) callCallbacks(c Cell) {
	// recursively call callbacks if cell was updated
	for _, cell := range m.deps[c] {
		if !cell.updated {
			continue
		}
		cell.updated = false // Ensure callbacks only called once
		for _, f := range cell.callbacks {
			f(cell.val)
		}
		m.callCallbacks(cell)
	}
}