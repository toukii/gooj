package utils

import (
// "strings"
)

func SubString(input string, length int) string {
	leng := len(input)
	if leng <= length {
		return input
	}
	return input[:length] + "..."
}
