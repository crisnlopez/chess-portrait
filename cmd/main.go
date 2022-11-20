package main

import (
	"errors"
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
	colorRed    = color.RGBA{255, 0, 0, 1}
)

func main() {
	fmt.Println("Init render")
	// Open Example Game
	file, err := os.Open("sample.pgn")
	if err != nil {
		log.Fatal(err)
	}

	// Create game
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

	// Render last move and marks squares
	lastMove := moveHistory[len(moveHistory)-1]

	whiteMarks, err := getMarkedSquares(moveHistory, 1)
	if err != nil {
		panic(err)
	}

	blackMarks, err := getMarkedSquares(moveHistory, 2)
	if err != nil {
		panic(err)
	}

	err = image.SVG(output, lastMove.PostPosition.Board(), image.MarkSquares(colorYellow, whiteMarks...), image.MarkSquares(colorRed, blackMarks...))
	if err != nil {
		log.Panicf("error image.SVG: %s", err)
	}
	fmt.Println("Finish program")
}

//printAll moves - util for debugging
func printMoves(moves []*chess.MoveHistory, writer io.Writer) error {
	for i := range moves {
		fmt.Printf("Board position %d: %s\n", i, moves[i].PostPosition.Board().Draw())
	}

	return nil
}

//getMarkedSquares return markedSquares for a given game
//sideBoard == 0 return all markedSquares
//sideBoard == 1 return whites marked Squares
//sideBoard == 2 return black marked Squares
//another number returns an error
func getMarkedSquares(moves []*chess.MoveHistory, sideBoard int) ([]chess.Square, error) {
	if sideBoard < 0 && 3 < sideBoard {
		return nil, errors.New("getMarkedSquares: only accept 0, 1 or 2")
	}

	markedSquares := make([]chess.Square, 0, len(moves))

	//mark white moves
	if sideBoard == 1 {
		for i := range moves {
			if i%2 != 0 {
				markedSquares = append(markedSquares, moves[i].Move.S1(), moves[i].Move.S2())
			}
		}

		return markedSquares, nil
	}

	//mark black moves
	if sideBoard == 2 {
		for i := range moves {
			if i%2 == 0 {
				markedSquares = append(markedSquares, moves[i].Move.S1(), moves[i].Move.S2())
			}
		}

		return markedSquares, nil
	}

	//mark all moves
	for i := range moves {
		markedSquares = append(markedSquares, moves[i].Move.S1(), moves[i].Move.S2())
	}

	return markedSquares, nil

}
