package main

// Chaser is a smart enemy! 0.0
type Chaser struct {
	Ghost
	path []string
}

var chaserPath []string

// NewChaser creates a new ghost
func NewChaser(row, col int, img string) *Chaser {
	var c Chaser
	c.position = Point{row, col}
	c.img = img
	return &c
}

func (c *Chaser) Pos() (row, col int) {
	row = c.position.row
	col = c.position.col
	return
}

func (c *Chaser) Img() string {
	return c.img
}

func (c *Chaser) Move() {
	dir := c.drawDirection()
	c.position = makeMove(c.position, dir)

	for _, s := range sprites {
		switch p := s.(type) {
		case *Player:
			if p.position == c.position {
				p.Kill()
				c.path = nil
			}
		}
	}
}

type Node struct {
	Point
	h int
	g int
}

func (c *Chaser) drawDirection() string {
	if len(c.path) == 0 {
		target := player.position
		c.path = c.find(c.position, target)
	}
	dir := c.path[0]
	c.path = c.path[1:]
	return dir
}

// implements A* pathfinding algorithm
//
// move cost = 1
// no diagonal movement allowed
// heuristic = manhattan distance
func (c *Chaser) find(origin Point, target Point) []string {
	var pf PathFinder
	path := pf.walk(origin, target)

	var directions []string
	current := origin
	for len(path) > 0 {
		next := path[len(path)-1]
		path = path[:len(path)-1]
		directions = append(directions, giveDirection(current, next))
		current = next
	}
	chaserPath = directions

	// return []string{c.Ghost.drawDirection()}
	return directions
}

func giveDirection(curr, dest Point) string {
	switch {
	case curr.row-dest.row == 0 && dest.col-curr.col == 1:
		return "RIGHT"
	case curr.row-dest.row == 0 && dest.col-curr.col == -1:
		return "LEFT"
	case curr.col-dest.col == 0 && dest.row-curr.row == 1:
		return "DOWN"
	case curr.col-dest.col == 0 && dest.row-curr.row == -1:
		return "UP"
	default:
		return "NOP"
	}
}

type PathFinder struct {
	open   PointSet
	closed PointSet
	table  map[Point]PointInfo
}

type PointSet map[Point]bool
type PointInfo struct {
	g      int
	h      int
	parent *Point
}

func (p PointInfo) Cost() int {
	return p.g + p.h
}

func (pf *PathFinder) nextPoint() *Point {
	var min *Point
	for k := range pf.open {
		if min == nil {
			min = &k
			continue
		}
		if pf.table[k].Cost() < pf.table[*min].Cost() {
			min = &k
		}
	}
	return min
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func distance(p1, p2 Point) int {
	return abs(p1.row-p2.row) + abs(p1.col-p2.col)
}

func (pf *PathFinder) isClosed(p Point) bool {
	_, ok := pf.closed[p]
	return ok
}

func (pf *PathFinder) isOpen(p Point) bool {
	_, ok := pf.open[p]
	return ok
}

func (pf *PathFinder) walk(start Point, target Point) []Point {
	var path []Point

	pf.open = make(PointSet)
	pf.closed = make(PointSet)
	pf.table = make(map[Point]PointInfo)

	pf.open[start] = true
	pf.table[start] = PointInfo{
		h: distance(start, target),
	}
	for current := pf.nextPoint(); current != nil; current = pf.nextPoint() {
		var neighbors []Point
		if up, err := current.Up(); err == nil {
			neighbors = append(neighbors, up)
		}
		if down, err := current.Down(); err == nil {
			neighbors = append(neighbors, down)
		}
		if left, err := current.Left(); err == nil {
			neighbors = append(neighbors, left)
		}
		if right, err := current.Right(); err == nil {
			neighbors = append(neighbors, right)
		}

		for _, n := range neighbors {
			switch {
			case n == target:
				// found path

				path = append(path, n)
				for pf.table[*current].parent != nil {
					path = append(path, *current)
					current = pf.table[*current].parent
				}
				return path
			case pf.isClosed(n):
				// do nothing
			case pf.isOpen(n):
				info := PointInfo{
					g:      pf.table[*current].g + 1,
					h:      distance(n, target),
					parent: current,
				}
				if pf.table[n].Cost() > info.Cost() {
					// found a better cost to n
					pf.table[n] = info
				}
			default:
				pf.open[n] = true
				pf.table[n] = PointInfo{
					g:      pf.table[*current].g + 1,
					h:      distance(n, target),
					parent: current,
				}
			}
		}
		// done with current
		// add to closed list
		// remove from open list
		pf.closed[*current] = true
		delete(pf.open, *current)
	}
	return path
}
