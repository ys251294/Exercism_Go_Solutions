package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	//panic("Please implement the CollatzConjecture function")
	steps := 0

	if n <= 0 {
		return 0, errors.New("error: n cannot be less than or zero")
	} else if n == 1 {
		return steps, nil
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
			steps++

		} else {
			n = n*3 + 1
			steps++
		}
	}

	return steps, nil
}
