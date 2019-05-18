package react

type holder struct {
	callback func(int)
	manager *callbackManager
}

func (h *holder) Cancel() {
	h.manager.remove(h)
}

type callbackManager struct {
	callbacks []*holder
}

func (m *callbackManager) AddCallback(f func(int)) Canceler {
	return m.add(f)
}

func (m *callbackManager) add(f func(int)) Canceler {
	h := &holder{f, m}
	m.callbacks = append(m.callbacks, h)
	return h
}

func (m *callbackManager) remove(h *holder) {
	for i := range m.callbacks {
		if m.callbacks[i] == h {
			n := len(m.callbacks)
			m.callbacks[i], m.callbacks[n - 1] = m.callbacks[n - 1], m.callbacks[i]
			m.callbacks = m.callbacks[:n - 1]
			break
		}
	}
}

func (m callbackManager) broadcast(val int) {
	for _, h := range m.callbacks {
		h.callback(val)
	}
}

type compute1Cell struct {
	cell
	callbackManager
	input   *Cell
	compute func(int) int
}

func (c *compute1Cell) update() {
	newValue := c.compute((*c.input).Value())
	if newValue != c.value {
		c.value = newValue
		c.notify()
		c.callbackManager.broadcast(newValue)
	}
}

type compute2Cell struct {
	cell
	callbackManager
	inputs  [2]*Cell
	compute func(int, int)int
}

func (c *compute2Cell) update() {
	newValue := c.compute((*c.inputs[0]).Value(), (*c.inputs[1]).Value())
	if newValue != c.value {
		c.value = newValue
		c.notify()
		c.callbackManager.broadcast(newValue)
	}
}
