package raindrops

import "fmt"

func Convert(number int) string {
	representation := ``
	if number % 3 == 0 {
		representation += `Pling`
	}
	if number % 5 == 0 {
		representation += `Plang`
	}
	if number % 7 == 0 {
		representation += `Plong`
	}
	if len(representation) == 0 {
		representation = fmt.Sprint(number)
	}
	return representation
}
