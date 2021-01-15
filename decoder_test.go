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
		// {"shift 1, but end of alphabet", 1, "xyzabc", "abcbcd"},
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
