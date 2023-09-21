package main

import (
	"fmt"
	"os"

	"github.com/dilyar85/chess/game"
)

func main() {

	fmt.Println("Enter the room ID")

	chessGame := game.New()

	arg := os.Args
	isInteractiveMode := len(arg) <= 1

	if isInteractiveMode {
		chessGame.StartInteractiveMode()
	} else {
		chessGame.StartFileMode(arg[1])
	}

}
