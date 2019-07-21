package main

import "errors"

// Sprite is any game character
type Sprite interface {
	Move()
	Pos() (int, int)
	Img() string
}

type Point struct {
	row, col int
}

func (p Point) Up() (Point, error) {
	p.row--
	if p.row < 0 {
		p.row = len(maze) - 1
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func (p Point) Down() (Point, error) {
	p.row++
	if p.row == len(maze) {
		p.row = 0
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func (p Point) Left() (Point, error) {
	p.col--
	if p.col < 0 {
		p.col = len(maze[0]) - 1
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func (p Point) Right() (Point, error) {
	p.col++
	if p.col == len(maze[0]) {
		p.col = 0
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func IsLegal(pos Point) bool {
	return maze[pos.row][pos.col] != '#'
}

func makeMove(oldPos Point, dir string) Point {
	var fn func() (Point, error)
	switch dir {
	case "UP":
		fn = oldPos.Up
	case "DOWN":
		fn = oldPos.Down
	case "LEFT":
		fn = oldPos.Left
	case "RIGHT":
		fn = oldPos.Right
	default:
		return oldPos
	}

	pos, err := fn()
	if err != nil {
		return oldPos
	}

	return pos
}
