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
		{"shift 1, but end of alphabet", 1, "xyzabc", "yzabcd"},
		{"shift 1, but end of alphabet", 1, "xYZaBc", "yZAbCd"},
		{"shift 1, but end of alphabet", 1, "1234567890", "1234567890"},
		{"shift 1, but end of alphabet", 1, "&é\"'(§è!çà)-", "&é\"'(§è!çà)-"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := newDecoder(tt.shift)
			output := d.decode(tt.input)
			if output != tt.expectedOutput {
				t.Errorf("excpected output %s with input %s and shift %d. But got %s", tt.expectedOutput, tt.input, tt.shift, output)
			}
		})
	}
}
