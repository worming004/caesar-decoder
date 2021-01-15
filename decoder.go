package main

import "container/ring"

type decoder interface {
	decode(input string) string
}
type decod struct {
	shiftRing *ring.Ring
	shift     int
}

func (d decod) decode(input string) string {
	result := ""
	for _, c := range input {
		newRune := c + rune(d.shift)
		result = result + string(newRune)
	}
	return result
}

func newDecoder(shift int) decoder {
	return decod{
		shift: shift,
	}
}
