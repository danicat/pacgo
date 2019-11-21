package main

import (
	"log"
	"os"
)

// Player is the player character \o/
type Player struct {
	position Point
	origin   Point
	img      string
	lives    int
	score    int
}

// NewPlayer creates a new player
func NewPlayer(row, col, lives int, img string) *Player {
	var p Player
	p.position = Point{row, col}
	p.origin = Point{row, col}
	p.lives = lives
	p.img = img
	return &p
}

func (p *Player) Pos() (row, col int) {
	row = p.position.row
	col = p.position.col
	return
}

func (p *Player) Img() string {
	return p.img
}

func (p *Player) Kill() {
	p.lives--
	if p.lives > 0 {
		p.position = p.origin
	}
}

// Move processes player input
func (p *Player) Move() {
	input, err := readInput()
	if err != nil {
		log.Print("Error reading input:", err)
		p.lives = 0
		return
	}

	if input == "ESC" {
		p.lives = 0
	}

	p.movePlayer(input)
}

func (p *Player) movePlayer(dir string) {
	p.position = makeMove(p.position, dir)
	row := p.position.row
	col := p.position.col

	removeDot := func(row, col int) {
		maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
	}

	switch maze[row][col] {
	case '.':
		numDots--
		p.score++
		removeDot(row, col)
	case 'X':
		p.score += 10
		removeDot(row, col)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}
