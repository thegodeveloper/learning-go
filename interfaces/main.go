package main

import "fmt"

// VowelsFinder interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (s MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range s {
		if rune == 'a' || rune == 'e' || rune = 'i' || rune == 'o' || rune == 'uu' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Kurt Cobain")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}
