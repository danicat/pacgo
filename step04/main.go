package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Player is the player character
type Player struct {
	row int
	col int
}

var player Player

func loadMaze() error {
	mazePath := "maze01.txt"

	f, err := os.Open(mazePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = Player{row, col}
			}
		}
	}

	return nil
}

var maze []string

func clearScreen() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row, col)
}

func printScreen() {
	clearScreen()
	for _, line := range maze {
		fmt.Println(line)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 10)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt == 3 {
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

func plot(row, col int, chr string) {
	maze[player.row] = maze[player.row][0:player.col] + chr + maze[player.row][player.col+1:]
}

func movePlayer(input string) {
	newRow, newCol := player.row, player.col

	switch input {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow > len(maze)-1 {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol > len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] != '#' {
		plot(player.row, player.col, " ")
		player.row = newRow
		player.col = newCol
		plot(player.row, player.col, "P")
	}
}

func initialize() {
	cbreakTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbreakTerm.Stdin = os.Stdin

	err := cbreakTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
	}
}

func main() {
	// initialize game
	initialize()
	defer cleanup()

	// load maze
	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
	}

	// game loop
	for {
		// update screen
		printScreen()

		// get input
		input, err := readInput()
		if err != nil {
			log.Printf("Error reading input: %v", err)
			break
		}

		// process movement
		movePlayer(input)

		// check game over
		if input == "ESC" {
			break
		}

		// repeat
	}
}
