package main

import (
	"fmt"
	"log"
	"os"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

func main() {
	fmt.Println("Init render")
	file, err := os.Open("sample.pgn")
	if err != nil {
		log.Fatal(err)
	}

	pgn, err := chess.PGN(file)
	if err != nil {
		log.Fatal(err)
	}
	game := chess.NewGame(pgn)

	positionHistory := game.Positions()
	lastPositionIdx := len(positionHistory) - 1
	lastPosition := positionHistory[lastPositionIdx]

	output, err := os.Create("output.svg")
	if err != nil {
		log.Fatal(err)
	}

	lasPositionData, err := lastPosition.MarshalText()
	if err != nil {
		panic(err)
	}

	var pos chess.Position
	pos.UnmarshalText(lasPositionData)
	if err := image.SVG(output, pos.Board()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finish program")
}
