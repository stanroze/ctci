package Chapter_1


// One Away: There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to
// check if they are one edit (or zero edits) away.
// EXAMPLE
// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false
func OneAway(str1, str2 string) bool {
	diff := len(str1) - len(str2)
	if diff < 0 {
		diff = diff * -1
	}

	if diff > 2 {
		return false
	}

	s1, s2 := []rune(str1), []rune(str2)
	var short, long []rune
	if len(s1) < len(s2) {
		short, long = s1, s2
	} else {
		short, long = s2, s1
	}

	si, li := 0, 0
	edit := false
	for si < len(short) && li < len(long) {
		if short[si] != long[li] {
			if edit {
				return false
			}
			edit = true

			// if long and short are the same length
			// then the short index gets to move ahead
			// because we are basically matching each char
			if len(short) == len(long) {
				si++
			}

		} else {
			// same letter in same position always increment
			si++
		}

		// always increment long index
		li++
	}

	return true
}
