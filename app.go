package main

import (
	"fmt"
	"os"
)

type app struct {
	alldecoder []decoder
}

func newApp() app {
	var alldecoder = make([]decoder, 26)
	for i := 0; i < len(alldecoder); i++ {
		alldecoder[i] = newCaesarDecoder(i)
	}
	return app{alldecoder: alldecoder}
}

func (app app) run() {
	toTranslate := os.Args[1]
	for i, decoder := range app.alldecoder {
		fmt.Printf("%d: %s\n", i, decoder.decode(toTranslate))
	}
}
