package forth

import (
	"fmt"
	"regexp"
	"strings"
)

var definitionName = regexp.MustCompile("^[[:alpha:][:punct:]]+$")

type processor interface {
	add(operator)
}

type interpreter struct {
	terms []operator
}

func (i *interpreter) add(term operator) {
	i.terms = append(i.terms, term)
}

type compiler struct {
	name  string
	terms []operator
}

func (c *compiler) add(term operator) {
	c.terms = append(c.terms, term)
}

func (c *compiler) missingName() bool {
	return len(c.name) == 0
}

func (c *compiler) compile() operator {
	terms := c.terms
	return func(stack *stack) {
		for _, term := range terms {
			term(stack)
		}
	}
}

func (c *compiler) setName(name string) {
	if !definitionName.MatchString(name) {
		panic(fmt.Errorf("forth: invalid definition name '%s'", name))
	}
	c.name = strings.ToLower(name)
}
