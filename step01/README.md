# Step 01: Input and Output

## Reading a File

```go
func loadMaze() error {
	f, err := os.Open("maze01.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

var maze []string

```

## Printing to the Screen

```go

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}
```

## Updating the game loop

```go
func main() {
	// initialize game

	// load maze
	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

	// game loop
	for {
		// process input

		// process movement

		// process collisions

		// update screen
		printScreen()

		// check game over

		// Temp: break infinite loop
		break

		// repeat
	}
}
```