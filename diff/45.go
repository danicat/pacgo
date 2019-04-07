6a7
> 	"math/rand"
18a20,27
> // Ghost is the enemy that chases the player :O
> type Ghost struct {
> 	row int
> 	col int
> }
> 
> var ghosts []*Ghost
> 
36a46,47
> 			case 'G':
> 				ghosts = append(ghosts, &Ghost{row, col})
71a83,87
> 	for _, g := range ghosts {
> 		moveCursor(g.row, g.col)
> 		fmt.Printf("G")
> 	}
> 
141a158,175
> func drawDirection() string {
> 	dir := rand.Intn(4)
> 	move := map[int]string{
> 		0: "UP",
> 		1: "DOWN",
> 		2: "RIGHT",
> 		3: "LEFT",
> 	}
> 	return move[dir]
> }
> 
> func moveGhosts() {
> 	for _, g := range ghosts {
> 		dir := drawDirection()
> 		g.row, g.col = makeMove(g.row, g.col, dir)
> 	}
> }
> 
184a219
> 		moveGhosts()
