package main

import (
	"log"

	"github.com/crisnlopez/chess-portrait/gameparser"
)

const inputPath = "./input/sample.pgn"

func main() {
	if err := gameparser.ParseFile(inputPath); err != nil {
		log.Fatal(err.Error())
	}
}
