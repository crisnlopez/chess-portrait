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
	game.Move(lastMovement.Move)
	pos := &chess.Position{}

	if err := pos.UnmarshalText([]byte(game.FEN())); err != nil {
		return err
	}

	// write board SVG to file
	writeFile, err := os.OpenFile("output.svg", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	yellow := color.RGBA{255, 255, 0, 1}
	mark := image.MarkSquares(yellow, chess.D2, chess.D4)
	if err := image.SVG(writeFile, pos.Board(), mark); err != nil {
		return err
	}

	return nil
}
