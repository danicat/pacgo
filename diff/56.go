47a48,49
> 			case '.':
> 				numDots++
55a58,60
> var score int
> var numDots int
> var lives = 1
71a77,78
> 				fallthrough
> 			case '.':
89c96
< 	fmt.Printf("Row %v Col %v", player.row, player.col)
---
> 	fmt.Printf("Score %v\nRow %v Col %v\n", score, player.row, player.col)
155a163,169
> 	switch maze[player.row][player.col] {
> 	case '.':
> 		numDots--
> 		score++
> 		// Remove dot from the maze
> 		maze[player.row] = maze[player.row][0:player.col] + " " + maze[player.row][player.col+1:]
> 	}
221a236,240
> 		for _, g := range ghosts {
> 			if player.row == g.row && player.col == g.col {
> 				lives = 0
> 			}
> 		}
227c246
< 		if input == "ESC" {
---
> 		if input == "ESC" || numDots == 0 || lives == 0 {
