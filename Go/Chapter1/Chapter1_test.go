package Chapter_1

import (
	"testing"
	"reflect"
)

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

func TestRotateMatrix(t *testing.T) {
	actual := [][]rune{
		{'1', '2', '3', '4', '5'},
		{'6', '7', '8', '9', 'a'},
		{'b', 'c', 'd', 'e', 'f'},
		{'g', 'h', 'i', 'j', 'k'},
		{'k', 'l', 'm', 'n', 'o'},
	}

	expected := [][]rune{
		{'5','a','f','k','o'},
		{'4','9','e','j','n'},
		{'3','8','d','i','m'},
		{'2','7','c','h','l'},
		{'1','6','b','g','k'},
	}

	result := RotateMatrix(actual)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RotateMatrix: Expected matrix didn't match the result matrix.")
	}

}

func TestZeroMatrix(t *testing.T) {
	actual := [][]int {
		{2,2,2,2},
		{2,2,0,2},
		{2,2,2,2},
	}

	expected := [][]int{
		{2, 2, 0, 2},
		{0, 0, 0, 0},
		{2,2,0,2},
	}

	result := ZeroMatrix(actual)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ZeroMatrix: failed, %v", result)
	}

	actual = [][]int {
		{2,0,2,2},
		{2,2,2,2},
		{2,2,2,2},
	}

	expected = [][]int{
		{0, 0, 0, 0},
		{2, 0, 2, 2},
		{2, 0, 2, 2},
	}

	result = ZeroMatrix(actual)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ZeroMatrix: failed, %v", result)
	}

	actual = [][]int {
		{2,0,2,2},
		{2,2,2,2},
		{2,2,0,2},
	}

	expected = [][]int{
		{0,0,0,0},
		{2,0,0,2},
		{0,0,0,0},
	}

	result = ZeroMatrix(actual)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ZeroMatrix: failed, %v", result)
	}


	actual = [][]int {
		{2,0,2,2},
		{0,2,2,2},
		{2,2,0,2},
	}

	expected = [][]int{
		{0,0,0,0},
		{0,0,0,0},
		{0,0,0,0},
	}

	result = ZeroMatrix(actual)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ZeroMatrix: failed, %v", result)
	}

	actual = [][]int {
		{2,2,2,2},
		{2,2,2,2},
		{2,2,2,2},
	}

	expected = [][]int{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}

	result = ZeroMatrix(actual)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ZeroMatrix: failed, %v", result)
	}
}

func TestIsRotation(t *testing.T) {
	s1 := "waterbottle"
	s2 := "bottlewater"

	if !IsRotation(s1, s2) {
		t.Errorf("IsRotation failed.")
	}
}