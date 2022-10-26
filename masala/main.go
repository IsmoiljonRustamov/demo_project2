package main

import (
	"unicode"
)

func checkPassword(pass string) bool {
	if len(pass) < 8 {
		return false
	}
	for _, r := range pass {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsUpper(r) && !unicode.IsLower(r) && !unicode.IsMark(r){
			return false
		}
	}
	return true
}
