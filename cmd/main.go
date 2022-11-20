package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

var (
	colorYellow = color.RGBA{255, 255, 0, 1}
)

func main() {
	fmt.Println("Init render")
	// Open Example Game
	file, err := os.Open("sample.pgn")
	if err != nil {
		log.Fatal(err)
	}

	// Read game
	pgn, err := chess.PGN(file)
	if err != nil {
		log.Fatal(err)
	}
	game := chess.NewGame(pgn)
	moveHistory := game.MoveHistory()

	// Create output file
	output, err := os.Create("output.svg")
	if err != nil {
		log.Fatal(err)
	}

	// Print moves
	fmt.Println("Printing moves")
	err = printMoves(moveHistory, output)
	if err != nil {
		panic(err)
	}

	// Render last move
	lastMove := moveHistory[len(moveHistory)-1]
	err = image.SVG(output, lastMove.PostPosition.Board(), image.MarkSquares(colorYellow, lastMove.Move.S1(), lastMove.Move.S2()))
	if err != nil {
		log.Panicf("error image.SVG: %s", err)
	}
	fmt.Println("Finish program")
}

func printMoves(moves []*chess.MoveHistory, writer io.Writer) error {
	for i := range moves {
		fmt.Printf("Board position %d: %s\n", i, moves[i].PostPosition.Board().Draw())
	}

	return nil
}
