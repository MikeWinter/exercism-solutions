package react

type inputCell struct {
	cell
}

func (c *inputCell) SetValue(val int) {
	if c.value != val {
		c.value = val
		c.notify()
	}
}
