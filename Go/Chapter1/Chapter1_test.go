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

func TestURLify(t *testing.T) {
	cases := []struct {
		str1 string
		length int
		expect string
	}{
		{"Mr John Smith    ", 13 ,"Mr%20John%20Smith"},
	}

	for _, tc := range cases {
		result := URLify(tc.str1, tc.length)
		if result != tc.expect {
			t.Errorf("CheckPermutation: input (%q, %d) expect %v, but got %v\n", tc.str1,  tc.length ,tc.expect, result)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		str1 string
		expect bool
	}{
		{"Tact Coa", true},
		{"abababa", true},
		{"babba", true},
		{"babbab", true},
		{"babbbabc", false},
	}

	for _, tc := range cases {
		result := IsPalindrome(tc.str1)
		if result != tc.expect {
			t.Errorf("IsPalindrome: input %q expect %v, but got %v\n", tc.str1,tc.expect, result)
		}
	}
}

func TestOneAway(t *testing.T) {
	cases := []struct {
		str1 string
		str2 string
		expect bool
	}{
		{"pale","ple", true},
		{"pales","pale", true},
		{"pale","bale", true},
		{"pale","bake", false},
		{"pale","pale", true},
		{"pale","plae", false},
	}

	for _, tc := range cases {
		result := OneAway(tc.str1, tc.str2)
		if result != tc.expect {
			t.Errorf("OneAway: input %q,%q expect %v, but got %v\n", tc.str1,tc.str2,tc.expect, result)
		}
	}
}

func TestCompress(t *testing.T) {

	cases := []struct {
		str1 string
		expect string
	}{
		{"aabcccccaaa","a2b1c5a3"},
		{"abcd","abcd"},
		{"aaabbccdd","a3b2c2d2"},
	}

	for _, tc := range cases {
		result := Compress(tc.str1)
		if result != tc.expect {
			t.Errorf("Compress: input %q expect %v, but got %v\n", tc.str1,tc.expect, result)
		}
	}
}