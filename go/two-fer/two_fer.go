package twofer

import "fmt"

func ShareWith(name string) string {
	var nameWithDefault string
	if name == "" {
		nameWithDefault = "you"
	} else {
		nameWithDefault = name
	}
	return fmt.Sprintf("One for %v, one for me.", nameWithDefault)
}
