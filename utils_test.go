package main

import (
	"testing"

	"github.com/iambighead/goutils/utils"
)

func TestStringArrayContains(t *testing.T) {
	testcases := []struct {
		input_arrays []string
		substring    string
		want         bool
	}{
		{[]string{"hello world", "something"}, "hello world", true},
		{[]string{"hello world", "something"}, "something", true},
		{[]string{"hello world", "something"}, "something1", false},
		{[]string{"hello world", "something"}, "hello", false},
		{[]string{"hello world", "something"}, "world", false},
	}

	for i, tc := range testcases {
		result := utils.StringArrayContains(tc.input_arrays, tc.substring)
		if result != tc.want {
			t.Errorf("StringArrayContains: test case #%d: %s: %t, want %t", i, tc.substring, result, tc.want)
		}
	}
}

func TestGetFileSha256InHex(t *testing.T) {
	testcases := []struct {
		inputfile string
		want      string
	}{
		{"testfile.txt", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{"testfile2.txt", "2d8c2f6d978ca21712b5f6de36c9d31fa8e96a4fa5d8ff8b0188dfb9e7c171bb"},
	}

	for i, tc := range testcases {
		result, _ := utils.GetFileSha256InHex(tc.inputfile)
		if string(result) != tc.want {
			t.Errorf("StringArrayContains: test case #%d: %s: %s, want %s", i, tc.inputfile, result, tc.want)
		}
	}
}
