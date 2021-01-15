package main

import (
	"container/ring"
	"fmt"
)

const runeToShiftLowerCase = "abcdefghijklmnopqrstuvwxyz"
const runeToShiftUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	rg := ring.New(len(runeToShift))
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

func (sr *shiftRing) getCharOnShift(shift int) rune {
	return sr.Move(shift).Value.(rune)
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
