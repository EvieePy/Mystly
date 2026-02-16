package core

import "unicode"

func ContainsSpace(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}

	return false
}
