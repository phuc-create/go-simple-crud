package helpers

import "unicode"

func ContainWhiteSpace(str string) bool {
	for _, c := range str {
		if unicode.IsSpace(c) {
			return true
		}
	}
	return false
}
