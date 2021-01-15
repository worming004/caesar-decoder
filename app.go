package main

import (
	"fmt"
	"os"
)

type app struct {
	alldecoders []decoder
}

func newApp() app {
	var alldecoders = make([]decoder, 26)
	for i := 0; i < len(alldecoders); i++ {
		alldecoders[i] = newCaesarDecoder(i)
	}
	return app{alldecoders: alldecoders}
}

func (app app) run() {
	toTranslate := os.Args[1]
	for i, decoder := range app.alldecoders {
		fmt.Printf("%d: %s\n", i, decoder.decode(toTranslate))
	}
}
