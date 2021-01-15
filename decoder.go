package main

import "container/ring"

type decoder interface {
	decode(input string) string
}
type caesarDecod struct {
	shiftRing *ring.Ring
	shift     int
}

func (d caesarDecod) decode(input string) string {
	result := ""
	for _, rToTrans := range input {
		lowerShift := newShiftRing(lower)
		upperShift := newShiftRing(upper)
		s, err := lowerShift.stopOnChar(rToTrans)
		if err == nil {
			result = result + string(s.getCharOnShift(d.shift))
			continue
		}
		s, err = upperShift.stopOnChar(rToTrans)
		if err == nil {
			result = result + string(s.getCharOnShift(d.shift))
			continue
		}
		result = result + string(rToTrans)
	}
	return result
}

func newCaesarDecoder(shift int) decoder {
	return caesarDecod{
		shift: shift,
	}
}
