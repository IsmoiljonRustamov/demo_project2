package main

import "testing"

func TestPass(t *testing.T) {
	res := checkPassword("!Isma02df") 
	if res != true {
		t.Errorf("Siz xato parol kiritingiz!")
	}
}