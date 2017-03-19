package Chapter_1

// URLi : Write a method to replace all spaces in a string with '%20
// You may assume that the string has suf client space at the end to hold the additional characters,
// and that you are given the "true" length of the string.
//
// EXAMPLE
// Input: "Mr John Smith ", 13
// Output: "Mr%20John%20Smith"
func URLify(url string, tl int) string {
	// start at the true length and go backwards
	// because Go doesn't let you modify strings in place
	// we have to convert this string to a []rune which
	// conveniently also covers Unicode

	urlRunes := []rune(url)
	i := len(url) - 1;
	for j := tl - 1;  j > 0; j-- {
		if urlRunes[j] == ' ' {
			urlRunes[i] = '0'
			urlRunes[i-1] = '2'
			urlRunes[i-2] = '%'
			i = i -3
		} else {
			urlRunes[i] = urlRunes[j]
			i = i - 1
		}
	}

	return string(urlRunes)
}
