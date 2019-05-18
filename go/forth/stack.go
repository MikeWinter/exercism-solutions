package forth

type stack []int

func (stack *stack) push(value int) {
	*stack = append(*stack, value)
}

func (stack *stack) pop() int {
	lastIndex := len(*stack) - 1
	value := (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]
	return value
}
