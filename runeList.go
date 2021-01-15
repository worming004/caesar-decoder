package main

import (
	"container/ring"
	"fmt"
)

const runeToShiftLowerCase = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const runeToShiftUpperCase = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type caseFilter int

const (
	lower caseFilter = 0
	upper caseFilter = 1
)

func newShiftRing(c caseFilter) shiftRing {
	var runeToShift string
	switch c {
	case lower:
		runeToShift = runeToShiftLowerCase
		break
	case upper:
		runeToShift = runeToShiftUpperCase
		break
	}
	rg := ring.New(52)
	length := rg.Len()
	for i := 0; i < length; i++ {
		rg.Value = rune(runeToShift[i])
		rg = rg.Next()
	}
	rg.Next()
	return shiftRing{rg}
}

type shiftRing struct {
	*ring.Ring
}

func (sr *shiftRing) stopOnChar(r rune) (*shiftRing, error) {
	if sr.Value.(rune) == r {
		return sr, nil
	}
	current := sr.Ring
	counter := 0
	for ; current.Value.(rune) != r; current = current.Next() {
		counter++
		if counter == sr.Len()+1 {
			return sr, fmt.Errorf("rune %v, %s not found", r, string(r))
		}
	}

	return &shiftRing{current}, nil
}
