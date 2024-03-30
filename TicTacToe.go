package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Define constants for players
const (
	PlayerX = "X"
	PlayerO = "O"
)

// Define the game board
var board [3][3]string

// Initialize the game board with empty cells
func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

// Print the game board
func printBoard() {
	clearScreen()
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -----")
		}
	}
}

// Clear the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Check if a player has won
func checkWin(player string) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||
			(board[0][i] == player && board[1][i] == player && board[2][i] == player) {
			return true
		}
	}
	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
		(board[0][2] == player && board[1][1] == player && board[2][0] == player) {
		return true
	}
	return false
}

// Check if the game board is full
func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

// Main function to play the game
func main() {
	initializeBoard()
	var currentPlayer = PlayerX
	var moveCount = 0
	for {
		printBoard()
		fmt.Printf("Player %s's turn (row column): ", currentPlayer)
		var row, col int
		_, err := fmt.Scan(&row, &col)
		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || board[row][col] != " " {
			fmt.Println("Invalid move. Try again.")
			continue
		}
		board[row][col] = currentPlayer
		moveCount++
		if checkWin(currentPlayer) {
			printBoard()
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		} else if isBoardFull() {
			printBoard()
			fmt.Println("It's a draw!")
			break
		}
		if currentPlayer == PlayerX {
			currentPlayer = PlayerO
		} else {
			currentPlayer = PlayerX
		}
	}
}
