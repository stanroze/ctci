package Chapter_1

import "testing"

func TestIsUniqueASCII(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"abcdefg", true},
		{"", true},
		{`_\/,"'.;`, true},
		{"12345", true},
		{"abca", false},
		{"aaaa", false},
		{"12341", false},
	}

	for _, thisCase := range cases {
		result := IsUniqueASCII(thisCase.input)
		if result != thisCase.expect {
			t.Errorf("IsUniqueASCII: input %q, expect %v, but got %v\n", thisCase.input, thisCase.expect, result)
		}
	}
}

func TestCheckPermutation(t *testing.T) {
	cases := []struct {
		str1 string
		str2 string
		expect bool
	}{
		{"abcdefg","gfedcba", true},
		{"","", true},
		{"aaabbb","ababab", true},
		{"abcdefg","abcd", false},
		{"abcd","aacd", false},
	}

	for _, tc := range cases {
		result := CheckPermutation(tc.str1, tc.str2)
		if result != tc.expect {
			t.Errorf("CheckPermutation: input %q, %q expect %v, but got %v\n", tc.str1,  tc.str2 ,tc.expect, result)
		}
	}
}
