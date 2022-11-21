package gameparser

import (
	"image/color"
	"os"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

func ParseFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	pgn, err := chess.PGN(f)
	if err != nil {
		return err
	}

	game := chess.NewGame(pgn)

	moveHistory := game.MoveHistory()

	lastMovement := moveHistory[len(moveHistory)-1]

	// write board SVG to file
	writeFile, err := os.Create("output.svg")
	if err != nil {
		return err
	}
	defer writeFile.Close()

	yellow := color.RGBA{255, 255, 0, 1}
	//TODO: marked position is not dynamic, but it should be
	mark := image.MarkSquares(yellow, lastMovement.Move.S1(), lastMovement.Move.S2())
	if err := image.SVG(writeFile, game.Position().Board(), mark); err != nil {
		return err
	}

	return nil
}
