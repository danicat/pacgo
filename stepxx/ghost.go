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
