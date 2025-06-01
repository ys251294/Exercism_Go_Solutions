package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	//panic("Implement the Distance function")
	if len(a) != len(b) {
		return -1, errors.New("strings are not the same length")
	}
	s1 := []rune(a)
	s2 := []rune(b)

	hammingDistance := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
