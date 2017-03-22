package Chapter_1

import (
	"unicode"
)


// Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
// EXAMPLE
// Input: Tact Coa
// Output: True (permutations: "taco cat", "atco eta", etc.)
func IsPalindrome(text string) bool  {
	count := make(map[rune]int,0)
	for _,c := range text {
		if c == ' ' {
			continue
		}

		count[unicode.ToLower(c)]++
	}
	alreadyHasOdd := false
	for _,v := range count {
		if v % 2 == 1 {
			if alreadyHasOdd {
				return false
			}
			alreadyHasOdd = true
		}
	}

	return true
}
