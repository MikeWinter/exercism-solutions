package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strings []string) FreqMap {
	c := make(chan FreqMap)
	m := FreqMap{}
	for _, s := range strings {
		go func(s string, c chan<- FreqMap) {
			c <- Frequency(s)
		}(s, c)
	}
	for i, n := 0, len(strings); i < n; i++ {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
