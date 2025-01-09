package encoder

import (
	"fmt"
	"testing"
)

func TestEncoder(t *testing.T) {
	var tests = []struct {
		plainText   string
		encodedText string
	}{
		{"encodethim", "ZW5jb2RldGhpbQ=="},
		{"alpha", "YWxwaGE="},
		{"a", "YQ=="},
		{"abcdef", "YWJjZGVm"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s", tt.plainText, tt.encodedText)
		t.Run(testname, func(t*testing.T){
			ans := EncodeString(tt.plainText)
			if ans != tt.encodedText {
				t.Errorf("got %s, want %s", ans, tt.encodedText)
			}
		})
	}
}

func TestDecoder(t *testing.T) {
	var tests = []struct {
		encodedText string
		plainText   string
	}{
		{"YWJjZGVm","abcdef"},
		{"YQ==","a"},
		{"YWxwaGE=","alpha"},
		{"ZW5jb2RldGhpbQ==","encodethim"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s", tt.encodedText, tt.plainText)
		t.Run(testname, func(t*testing.T){
			ans, _ := DecodeString(tt.encodedText)
			fmt.Println(len(ans), len(tt.plainText))
			if ans != tt.plainText {
				t.Errorf("got %s, want %s", ans, tt.plainText)
			}
		})
	}

	
}