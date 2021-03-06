// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

/*
~ TICTACTOE GAME IN GO ~
+ This example uses the very basics of the Go language.
+ The goal is learning all the basics.
*/

const maxTurns = 9

var (
	won    bool   // is there any winner?
	turn   int    // total valid turns played
	player string // current player

	cells     [maxTurns]string // used to draw the board: contains the players' moves
	lastPos   int              // last played position
	wrongMove bool             // was the last move wrong?
)

func main() {
	player = player1
	printBoard()

	for i := 0; i < 4; i++ {
		wait()
		play()
		printBoard()

		fmt.Printf("\n>>> PLAYER %q PLAYS to %d\n", player, lastPos+1)
		printEnding()
	}
}

func wait() {
	fmt.Println()
	fmt.Scanln()
}

// printStatus prints the current status of the game
// it cannot access to the names (vars, consts, etc) inside any other func
func printStatus() {
	fmt.Println()

	progress := (1 - (float64(turn) / maxTurns)) * 100
	fmt.Printf("Current Turn           : %d\n", turn)
	fmt.Printf("Is there a winner      : %t\n", won)
	fmt.Printf("Turns left             : %.1f%%\n", progress)
}

func printEnding() {
	switch {
	default:
		switchPlayer()
		printStatus()

	case wrongMove:
		fmt.Printf(">>> CELL IS OCCUPIED: PLAY AGAIN!\n")
		wrongMove = false // reset for the next turn

	case won:
		fmt.Printf(">>> WINNER: %s\n", player)

	case turn == maxTurns:
		fmt.Println(">>> TIE!")
	}
}
