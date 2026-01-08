// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
