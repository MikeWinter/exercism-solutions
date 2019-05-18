package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("lengths do not match")
	}

	length := len(a)
	distance := 0
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
