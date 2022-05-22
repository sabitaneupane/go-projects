package main

import "fmt"

const (
	width     = 50
	height    = 10
	cellEmpty = ' '
	cellBall  = '⚽'
)

var cell rune

func main() {
	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	buf := make([]rune, 0, width*height)

	board[0][0] = true
	buf = buf[:0]

	// draw the board
	for y := range board[0] {

		for x := range board {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			buf = append(buf, cell, ' ')
		}
		buf = append(buf, '\n')
	}
	fmt.Print(string(buf))
}
