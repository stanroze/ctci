package Chapter_1

import (
	"bytes"
	"strconv"
)


// String Compression: Implement a method to perform basic string compression using
// the counts of repeated characters. For example, the string aabcccccaaa would become a2blc5a3.
// If the "compressed" string would not become smaller than the original string, your method
// should return the original string. You can assume the string has only uppercase
// and lowercase letters (a - z).
func Compress(text string) string {
	if len(text) < 2 {
		return text
	}

	// would be different for ASCII set only
	// could count how many compressions needed and make a fix-sized array
	input := []rune(text)
	var compressed bytes.Buffer
	last := input[0]
	count := 1
	//totalLen :=0
	for i:=1; i<len(input);i++ {
		if input[i] == last {
			count++
		} else {
			compressed.WriteRune(last)
			compressed.WriteString(strconv.Itoa(count))
			last = input[i]
			count = 1
		}
	}

	compressed.WriteRune(last)
	compressed.WriteString(strconv.Itoa(count))

	// this is not ideal - better to count ahead of compression
	// to see if the compressed will be longer than uncompressed
	if compressed.Len() > len(input) {
		return text
	}

	return compressed.String()
}