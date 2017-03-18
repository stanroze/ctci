package Chapter_1

// Check Permutation: Given two strings,write a method to decide if one is a permutation of the
// other.
// Assumption: ASCII - for non-ascii use a map[rune]int
func CheckPermutation(str1, str2 string) bool {
	// if the len of 2 strings is different, then it's not a permutation
	if len(str1) != len(str2) {
		return false
	}

	var set [128]int
	for _,r := range str1 {
		set[r]++
	}

	for _,r := range str2 {
		if set[r] == 0 {
			return false
		}
		set[r]--
	}
	return true
}
