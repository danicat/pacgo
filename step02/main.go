package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}

	return "", nil
}

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
		input, err := readInput()
		if err != nil {
			log.Printf("Error reading input: %v", err)
			break
		}

		// process movement

		// process collisions

		// update screen
		printScreen()

		// check game over
		if input == "ESC" {
			break
		}

		// repeat
	}
}
