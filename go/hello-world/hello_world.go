package greeting

import "fmt"

func HelloWorld() string {
	a := factory(1)
	b := factory(2)

	fmt.Println(b(), a())

	return "Hello, World!"
}

func factory(v int) func() int {
	return func() int {
		return v
	}
}
