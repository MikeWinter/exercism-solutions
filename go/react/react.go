package react

type reactor struct{}

func New() Reactor {
	return reactor{}
}

type cell struct {
	value     int
	observers []*observer
}

func (c cell) Value() int {
	return c.value
}

type observer interface {
	Cell
	update()
}

type notifier interface {
	register(*observer)
	notify()
}

func (c *cell) register(observer *observer) {
	c.observers = append(c.observers, observer)
}

func (c cell) notify() {
	for _, o := range c.observers {
		(*o).update()
	}
}

func (reactor) CreateInput(val int) InputCell {
	return &inputCell{cell{value: val}}
}

func (reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	newCell := observer(&compute1Cell{cell: cell{value: f(c.Value())}, input: &c, compute: f})
	c.(notifier).register(&newCell)
	return newCell.(ComputeCell)
}

func (reactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	newCell := observer(&compute2Cell{cell: cell{value: f(c1.Value(), c2.Value())}, inputs: [2]*Cell{&c1, &c2}, compute: f})
	c1.(notifier).register(&newCell)
	c2.(notifier).register(&newCell)
	return newCell.(ComputeCell)
}
