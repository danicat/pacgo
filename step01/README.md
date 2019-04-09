# Step 01: Input and Output

In this lesson you will learn how to:

- Read from a file
- Print to the standard output
- Handle multiple return values
- Handle errors
- Create and add an element to a slice
- Iterate over a collection
- Defer a function call

## Overview

We've got the basics covered, now it's time to get this game started!

First, we are going to read the maze data. We have a file called `maze01.txt` that's basically an ASCII representation of the maze (you can open it in a text editor if you like). You may assume that:

- # represents a wall
- . represents a dot
- P represents the player
- G represents the ghosts (enemies)
- X represents the power up pills

Our first task consists in loading this ASCII representation of the maze to a slice of strings and then printing it to the screen. Looks simple, right? It is indeed!

## Task 01: Load the maze

Let's start by reading the `maze01.txt` file.

We are going to use the function `Open` from the `os` package to open the it, and an scanner object from the buffered IO package (`bufio`) to read it to memory (to a global variable called `maze`). Finally we need to release the file handler by calling `os.Close`. 

All that comes together as the code below:

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

Now let's break it down and see what's going on.

Please note that you need import both the `os` and `bufio` packages as shown below:

```go
import "os"
import "bufio"
```

Alternatively, since you already have one import (`fmt`), you can add it as a list:

```go
import (
    "bufio"
    "fmt"
    "os"
)
```

The `os.Open()` function return a pair of values: a file and an error. Returning multiple values from a function is a common pattern in Go, specially for returning errors.

```go
f, err := os.Open("maze01.txt")
```

The `:=` operator is an assignment operator, but with special property that it automatically infers the type of the variable(s) based on the value(s) on the right side. 

Keep in mind that Go is a strongly typed language, but that nice feature saves us the trouble of specifying the type when it's possible to infer it.

In the case above, Go automatically infers the type for both `f` and `err` variables.

When a function returns an error is a common patter to check it immediately afterwards:

```go
	f, err := os.Open("maze01.txt")
	if err != nil {
		// do something with err
	}
```

`nil` in Go means no value is assigned to a variable. 

The `if` statement executes a statement if the condition is true. It can optionally have an initialization clause just like the `for` statement, and an `else` clause that runs if the condition is false. Please keep in mind that the scope of the variable created will just be the if statement body. For example:

```go
// optional initialization clause
if foo := rand.Intn(2); foo == 0 {
    fmt.Print(foo) // foo is valid here
} else {
    fmt.Print(foo) // and here
}
// but can't use foo here!
```


## Task 02: Printing to the Screen

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

	// load resources
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