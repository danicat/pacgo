package main

import "math/rand"

// Ghost is the enemy that chases the player :O
type Ghost struct {
	position Point
	img      string
}

// NewGhost creates a new ghost
func NewGhost(row, col int, img string) *Ghost {
	var g Ghost
	g.position = Point{row, col}
	g.img = img
	return &g
}

func (g *Ghost) Pos() (row, col int) {
	row = g.position.row
	col = g.position.col
	return
}

func (g *Ghost) Img() string {
	return g.img
}

func (g *Ghost) Move() {
	dir := g.drawDirection()
	g.position = makeMove(g.position, dir)

	for _, s := range sprites {
		switch p := s.(type) {
		case *Player:
			if p.position == g.position {
				p.Kill()
			}
		}
	}
}

func (g *Ghost) drawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}
