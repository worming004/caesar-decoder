package main

import (
	"testing"
)

func TestDecoder(t *testing.T) {

	tests := []struct {
		name                  string
		shift                 int
		input, expectedOutput string
	}{
		{"shift 0", 0, "abc", "abc"},
		{"shift 1", 1, "abc", "bcd"},
		{"shift 5", 5, "abc", "fgh"},
		{"shift 5 uppercase", 5, "ABC", "FGH"},
		{"shift 3 uppercase", 3, "ABC", "DEF"},
		{"shift 1 uppercase", 1, "ABC", "BCD"},
		{"shift 0 uppercase", 0, "ABC", "ABC"},
		{"shift 1, but end of alphabet", 1, "xyzabc", "yzabcd"},
		{"shift 1, but end of alphabet", 1, "xYZaBc", "yZAbCd"},
		{"shift 1, but end of alphabet", 1, "1234567890", "1234567890"},
		{"shift 1, but end of alphabet", 1, "&é\"'(§è!çà)-", "&é\"'(§è!çà)-"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := newCaesarDecoder(tt.shift)
			output := d.decode(tt.input)
			if output != tt.expectedOutput {
				t.Errorf("expected output \n%s\n with input \n%s\n and shift %d. But got \n%s\n", tt.expectedOutput, tt.input, tt.shift, output)
			}
		})
	}
}
