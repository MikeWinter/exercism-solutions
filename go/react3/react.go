package react

func New() Reactor {
	return react{}
}

type react struct {}

func (react) CreateInput(value int) InputCell {
	return &input{cell: cell{value: value}}
}

func (react) CreateCompute1(in Cell, f func(int) int) ComputeCell {
	c := &compute1{compute{cell: cell{value: f(in.Value())}, in: in}, f}
	in.(node).add(c)
	return c
}

func (react) CreateCompute2(in1 Cell, in2 Cell, f func(int, int) int) ComputeCell {
	c := &compute2{compute{cell: cell{value: f(in1.Value(), in2.Value())}, in: in1}, in2, f}
	in1.(node).add(c)
	in2.(node).add(c)
	return c
}

type node interface {
	update(bubbling bool)
	children() []node
	add(node)
}

type cell struct {
	value int
	dependents []node
}

func (c cell) Value() int {
	return c.value
}

func (c *cell) add(dependent node) {
	c.dependents = append(c.dependents, dependent)
}

func (c *cell) children() []node {
	return c.dependents
}

type input struct {
	cell
}

func (i *input) SetValue(value int) {
	if i.value != value {
		i.value = value

		i.update(false)
	}
}

func (i *input) update(bool) {
	var visited []node
	pending := i.children()

	for j := 0; j < len(pending); j++ {
		n := pending[j]
		pending = append(pending, n.children()...)
		n.update(false)
		visited = append(visited, n)
	}

	for _, n := range visited {
		n.update(true)
	}
}

type holder struct {
	compute *compute
	cb func(int)
}

func (h *holder) Cancel() {
	h.compute.removeCallback(h)
}

type compute struct {
	cell
	in Cell
	callbacks []*holder
	dirty bool
}

func (c *compute) AddCallback(cb func(int)) Canceler {
	h := &holder{c, cb}
	c.callbacks = append(c.callbacks, h)
	return h
}

func (c *compute) removeCallback(h *holder) {
	for i := range c.callbacks {
		if c.callbacks[i] == h {
			lastIndex := len(c.callbacks) - 1
			c.callbacks[i] = c.callbacks[lastIndex]
			c.callbacks[lastIndex] = nil
			c.callbacks = c.callbacks[:lastIndex]
			break
		}
	}
}

func (c *compute) notifyCallbacks() {
	for _, h := range c.callbacks {
		h.cb(c.Value())
	}
}

type compute1 struct {
	compute
	cb func(int)int
}

func (c *compute1) update(bubbling bool) {
	if !bubbling {
		newValue := c.cb(c.in.Value())
		if newValue != c.value {
			c.value = newValue
			c.dirty = true
		}
	} else if c.dirty {
		c.notifyCallbacks()
		c.dirty = false
	}
}

type compute2 struct {
	compute
	in2 Cell
	cb func(int, int)int
}

func (c *compute2) update(bubbling bool) {
	if !bubbling {
		newValue := c.cb(c.in.Value(), c.in2.Value())
		if newValue != c.value {
			c.value = newValue
			c.dirty = true
		}
	} else if c.dirty {
		c.notifyCallbacks()
		c.dirty = false
	}
}
