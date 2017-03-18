package Chapter_1

// Is Unique: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?
//
// Assuming an ASCII only set (hint: only 128 chars)
func IsUniqueASCII(text string) bool {
	if len(text) > 128 {
		return false
	}

	var set [128]bool
	for _,c := range text {
		if set[c] {
			return false
		}

		set[c] = true
	}

	return true
}
