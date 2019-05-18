package forth

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type (
	dictionary map[string]operator
	operator   func(*stack)
	token      = string
)

var number = regexp.MustCompile("^0|[1-9][0-9]*$")

func Forth(expressions []string) (_ []int, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			switch recovered := recovered.(type) {
			case error:
				err = recovered
			default:
				err = fmt.Errorf("forth: unexpected error: %v", recovered)
			}
		}
	}()

	operators := dictionary{
		"+":    add,
		"-":    subtract,
		"*":    multiply,
		"/":    divide,
		"dup":  duplicate,
		"drop": drop,
		"swap": swap,
		"over": over,
	}
	terms := compile(expressions, operators)
	return evaluate(terms), nil
}

func evaluate(terms []operator) []int {
	var stack stack
	for _, term := range terms {
		term(&stack)
	}
	return stack
}

func compile(expressions []string, operators dictionary) []operator {
	terms := make([]operator, 0)
	for _, expression := range expressions {
		terms = append(terms, parse(expression, operators)...)
	}
	return terms
}

func parse(expression string, operators dictionary) []operator {
	var (
		interpreter interpreter
		complr      *compiler
	)

	var processor processor = &interpreter

	tokens := strings.Fields(expression)
	for _, token := range tokens {
		switch {
		case token == ":":
			complr = new(compiler)
			processor = complr
		case token == ";":
			operators[complr.name] = complr.compile()
			processor = &interpreter
			complr = nil
		case processor == complr && complr.missingName():
			complr.setName(token)
		case operators.contains(token):
			processor.add(operators.get(token))
		case number.MatchString(token):
			processor.add(literal(token))
		default:
			panic(fmt.Errorf("forth: unrecognised token '%s'", token))
		}
	}
	return interpreter.terms
}

func (dict dictionary) get(key string) operator {
	return dict[strings.ToLower(key)]
}

func (dict dictionary) contains(key string) bool {
	_, ok := dict[strings.ToLower(key)]
	return ok
}

func literal(token token) operator {
	val, _ := strconv.Atoi(token)
	return func(stack *stack) {
		stack.push(val)
	}
}

func add(stack *stack) {
	augend := stack.pop()
	addend := stack.pop()
	sum := augend + addend
	stack.push(sum)
}

func subtract(stack *stack) {
	subtrahend := stack.pop()
	minuend := stack.pop()
	difference := minuend - subtrahend
	stack.push(difference)
}

func multiply(stack *stack) {
	multiplicand := stack.pop()
	multiplier := stack.pop()
	product := multiplicand * multiplier
	stack.push(product)
}

func divide(stack *stack) {
	denominator := stack.pop()
	numerator := stack.pop()
	quotient := numerator / denominator
	stack.push(quotient)
}

func duplicate(stack *stack) {
	value := stack.pop()
	stack.push(value)
	stack.push(value)
}

func drop(stack *stack) {
	stack.pop()
}

func swap(stack *stack) {
	first := stack.pop()
	second := stack.pop()
	stack.push(first)
	stack.push(second)
}

func over(stack *stack) {
	first := stack.pop()
	second := stack.pop()
	stack.push(second)
	stack.push(first)
	stack.push(second)
}
