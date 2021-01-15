package main

import (
	"reflect"
	"testing"
)

func Test_shiftRing_stopOnChar(t *testing.T) {

	type args struct {
		r rune
	}
	tests := []struct {
		name    string
		sr      shiftRing
		args    args
		want    rune
		wantErr bool
	}{
		{"should find a", newShiftRing(lower), args{'a'}, 'a', false},
		{"should find b", newShiftRing(lower), args{'b'}, 'b', false},
		{"should find Z", newShiftRing(upper), args{'Z'}, 'Z', false},
		{"should find A", newShiftRing(upper), args{'A'}, 'A', false},
		{"should find z", newShiftRing(lower), args{'z'}, 'z', false},
		{"should find M", newShiftRing(upper), args{'M'}, 'M', false},
		{"should not find =", newShiftRing(upper), args{'='}, '=', true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.sr.stopOnChar(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("shiftRing.stopOnChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("shiftRing.stopOnChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
