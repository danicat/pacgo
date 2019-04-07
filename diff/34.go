10a11,18
> // Player is the player character \o/
> type Player struct {
> 	row int
> 	col int
> }
> 
> var player Player
> 
23a32,40
> 	for row, line := range maze {
> 		for col, char := range line {
> 			switch char {
> 			case 'P':
> 				player = Player{row, col}
> 			}
> 		}
> 	}
> 
41c58,66
< 		fmt.Println(line)
---
> 		for _, chr := range line {
> 			switch chr {
> 			case '#':
> 				fmt.Printf("%c", chr)
> 			default:
> 				fmt.Printf(" ")
> 			}
> 		}
> 		fmt.Printf("\n")
42a68,73
> 
> 	moveCursor(player.row, player.col)
> 	fmt.Printf("P")
> 
> 	moveCursor(len(maze)+1, 0)
> 	fmt.Printf("Row %v Col %v", player.row, player.col)
54a86,98
> 	} else if cnt >= 3 {
> 		if buffer[0] == 0x1b && buffer[1] == '[' {
> 			switch buffer[2] {
> 			case 'A':
> 				return "UP", nil
> 			case 'B':
> 				return "DOWN", nil
> 			case 'C':
> 				return "RIGHT", nil
> 			case 'D':
> 				return "LEFT", nil
> 			}
> 		}
59a104,141
> func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
> 	newRow, newCol = oldRow, oldCol
> 
> 	switch dir {
> 	case "UP":
> 		newRow = newRow - 1
> 		if newRow < 0 {
> 			newRow = len(maze) - 1
> 		}
> 	case "DOWN":
> 		newRow = newRow + 1
> 		if newRow == len(maze)-1 {
> 			newRow = 0
> 		}
> 	case "RIGHT":
> 		newCol = newCol + 1
> 		if newCol == len(maze[0]) {
> 			newCol = 0
> 		}
> 	case "LEFT":
> 		newCol = newCol - 1
> 		if newCol < 0 {
> 			newCol = len(maze[0]) - 1
> 		}
> 	}
> 
> 	if maze[newRow][newCol] == '#' {
> 		newRow = oldRow
> 		newCol = oldCol
> 	}
> 
> 	return
> }
> 
> func movePlayer(dir string) {
> 	player.row, player.col = makeMove(player.row, player.col, dir)
> }
> 
101a184
> 		movePlayer(input)
