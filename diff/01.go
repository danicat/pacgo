3c3,32
< import "fmt"
---
> import (
> 	"bufio"
> 	"fmt"
> 	"log"
> 	"os"
> )
> 
> func loadMaze() error {
> 	f, err := os.Open("maze01.txt")
> 	if err != nil {
> 		return err
> 	}
> 	defer f.Close()
> 
> 	scanner := bufio.NewScanner(f)
> 	for scanner.Scan() {
> 		line := scanner.Text()
> 		maze = append(maze, line)
> 	}
> 
> 	return nil
> }
> 
> var maze []string
> 
> func printScreen() {
> 	for _, line := range maze {
> 		fmt.Println(line)
> 	}
> }
8a38,42
> 	err := loadMaze()
> 	if err != nil {
> 		log.Printf("Error loading maze: %v\n", err)
> 		return
> 	}
24d57
< 		fmt.Println("Hello, Pac Go!")
