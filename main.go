package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	GRID_WIDTH  = 4
	GRID_HEIGHT = 4
)

type Board struct {
	grid  [GRID_WIDTH][GRID_HEIGHT]int
	score int
	over  bool
}

func NewBoard() *Board {
	board := &Board{}
	rand.Seed(time.Now().UnixNano())
	board.addRandomTile()
	board.addRandomTile()
	return board
}

func (b *Board) addRandomTile() {
	var emptyCells [][2]int

	for i := 0; i < GRID_WIDTH; i++ {
		for j := 0; j < GRID_HEIGHT; j++ {
			if b.grid[i][j] == 0 {
				emptyCells = append(emptyCells, [2]int{i, j})
			}
		}
	}

	if len(emptyCells) == 0 {
		b.over = true
		return
	}

	pos := emptyCells[rand.Intn(len(emptyCells))]
	if rand.Intn(10) < 9 { // 90% chance for 2, 10% for 4
		b.grid[pos[0]][pos[1]] = 2
	} else {
		b.grid[pos[0]][pos[1]] = 4
	}
}

func (b *Board) PrintBoard() {
	fmt.Println("Score:", b.score)
	for i := 0; i < GRID_HEIGHT; i++ {
		for j := 0; j < GRID_WIDTH; j++ {
			fmt.Printf("%4d ", b.grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) MoveUp() bool {
	moved := false
	for j := 0; j < GRID_WIDTH; j++ {
		for i := 1; i < GRID_HEIGHT; i++ {
			if b.grid[i][j] != 0 {
				k := i
				for k > 0 && b.grid[k-1][j] == 0 {
					b.grid[k-1][j] = b.grid[k][j]
					b.grid[k][j] = 0
					k--
					moved = true
				}
				if k > 0 && b.grid[k-1][j] == b.grid[k][j] {
					b.grid[k-1][j] *= 2
					b.score += b.grid[k-1][j]
					b.grid[k][j] = 0
					moved = true
				}
			}
		}
	}
	return moved
}

func (b *Board) MoveDown() bool {
	moved := false
	for j := 0; j < GRID_WIDTH; j++ {
		for i := GRID_HEIGHT - 2; i >= 0; i-- {
			if b.grid[i][j] != 0 {
				k := i
				for k < GRID_HEIGHT-1 && b.grid[k+1][j] == 0 {
					b.grid[k+1][j] = b.grid[k][j]
					b.grid[k][j] = 0
					k++
					moved = true
				}
				if k < GRID_HEIGHT-1 && b.grid[k+1][j] == b.grid[k][j] {
					b.grid[k+1][j] *= 2
					b.score += b.grid[k+1][j]
					b.grid[k][j] = 0
					moved = true
				}
			}
		}
	}
	return moved
}

func (b *Board) MoveLeft() bool {
	moved := false
	for i := 0; i < GRID_HEIGHT; i++ {
		for j := 1; j < GRID_WIDTH; j++ {
			if b.grid[i][j] != 0 {
				k := j
				for k > 0 && b.grid[i][k-1] == 0 {
					b.grid[i][k-1] = b.grid[i][k]
					b.grid[i][k] = 0
					k--
					moved = true
				}
				if k > 0 && b.grid[i][k-1] == b.grid[i][k] {
					b.grid[i][k-1] *= 2
					b.score += b.grid[i][k-1]
					b.grid[i][k] = 0
					moved = true
				}
			}
		}
	}
	return moved
}

func (b *Board) MoveRight() bool {
	moved := false
	for i := 0; i < GRID_HEIGHT; i++ {
		for j := GRID_WIDTH - 2; j >= 0; j-- {
			if b.grid[i][j] != 0 {
				k := j
				for k < GRID_WIDTH-1 && b.grid[i][k+1] == 0 {
					b.grid[i][k+1] = b.grid[i][k]
					b.grid[i][k] = 0
					k++
					moved = true
				}
				if k < GRID_WIDTH-1 && b.grid[i][k+1] == b.grid[i][k] {
					b.grid[i][k+1] *= 2
					b.score += b.grid[i][k+1]
					b.grid[i][k] = 0
					moved = true
				}
			}
		}
	}
	return moved
}

func main() {
	b := NewBoard()
	b.PrintBoard()

	for !b.over {
		var move string
		fmt.Scan(&move)
		moved := false

		switch move {
		case "w":
			moved = b.MoveUp()
		case "s":
			moved = b.MoveDown()
		case "a":
			moved = b.MoveLeft()
		case "d":
			moved = b.MoveRight()
		case "q":
			fmt.Println("Game quit")
			return
		default:
			fmt.Println("Use w/a/s/d to move, q to quit")
			continue
		}

		if moved {
			b.addRandomTile()
		}
		b.PrintBoard()
	}
	fmt.Println("Game Over! Final Score:", b.score)
}
