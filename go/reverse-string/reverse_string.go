package reverse

func String(str string) string {
	letters := []rune(str)
	length := len(letters)
	depth := length / 2

	for i := 0; i < depth; i++ {
		swap(letters, i, length-i-1)
	}
	return string(letters)
}

func swap(letters []rune, i int, j int) {
	letters[i], letters[j] = letters[j], letters[i]
}
