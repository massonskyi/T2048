package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
	if rand.Intn(10) < 9 {
		b.grid[pos[0]][pos[1]] = 2
	} else {
		b.grid[pos[0]][pos[1]] = 4
	}
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
	app := tview.NewApplication()
	board := NewBoard()

	// Create text view for displaying the game board
	gridView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(false)

	// Define color function before using it
	getColor := func(val int) string {
		switch val {
		case 2:
			return "cyan"
		case 4:
			return "blue"
		case 8:
			return "green"
		case 16:
			return "yellow"
		case 32:
			return "orange"
		case 64:
			return "red"
		case 128, 256, 512, 1024, 2048:
			return "purple"
		default:
			return "white"
		}
	}

	// Update display function
	updateDisplay := func() {
		var display string
		display += fmt.Sprintf("[yellow]Score: %d[white]\n\n", board.score)
		for i := 0; i < GRID_HEIGHT; i++ {
			for j := 0; j < GRID_WIDTH; j++ {
				val := board.grid[i][j]
				if val == 0 {
					display += "[gray]    -[white] "
				} else {
					display += fmt.Sprintf("[%s]%4d[white] ", getColor(val), val)
				}
			}
			display += "\n\n"
		}
		if board.over {
			display += "[red]Game Over![white]\n"
		}
		display += "[green]Use arrow keys to move, 'q' to quit[white]"
		gridView.SetText(display)
	}

	// Input handling
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if board.over {
			if event.Rune() == 'q' {
				app.Stop()
			}
			return event
		}

		moved := false
		switch event.Key() {
		case tcell.KeyUp:
			moved = board.MoveUp()
		case tcell.KeyDown:
			moved = board.MoveDown()
		case tcell.KeyLeft:
			moved = board.MoveLeft()
		case tcell.KeyRight:
			moved = board.MoveRight()
		case tcell.KeyRune:
			if event.Rune() == 'q' {
				app.Stop()
			}
		}
		if moved {
			board.addRandomTile()
			updateDisplay()
		}
		return event
	})

	// Initialize display
	updateDisplay()

	// Run application
	if err := app.SetRoot(gridView, true).Run(); err != nil {
		panic(err)
	}
}
