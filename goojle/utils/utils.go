package utils

import (
// "strings"
)

func SubString(input string, length int) string {
	leng := len(input)
	if leng <= length {
		return input
	}
	var i int
	for i, _ = range input {
		if i >= length {
			break
		}
	}
	return input[:i] + " ..."
}
