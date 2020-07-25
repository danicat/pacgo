# Step 01: Input and Output

In this lesson you will learn how to:

- Read from a file
- Print to the standard output
- Handle multiple return values
- Handle errors
- Create and add an element to a slice
- Range loop over a slice
- Defer a function call
- Log errors

## Overview

We've got the basics covered, now it's time to get this game started!

First, we are going to read the maze data. We have a file called `maze01.txt` that's basically an ASCII representation of the maze (you can open it in a text editor if you like). You may assume that:

```
- # represents a wall
- . represents a dot
- P represents the player
- G represents the ghosts (enemies)
- X represents the power up pills
```

Our first task consists of loading this ASCII representation of the maze to a slice of strings and then printing it to the screen. Looks simple, right? It is indeed!

## Task 01: Load the Maze

Let's start by reading the `maze01.txt` file.

We are going to use the function `Open` from the `os` package to open it, and a scanner object from the buffered IO package (`bufio`) to read it to memory (to a global variable called `maze`). Finally we need to release the file handler by calling `os.Close`.

All that comes together as the code below:

```go
var maze []string

func loadMaze(file string) error {
    f, err := os.Open(file)
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
```

Now let's break it down and see what's going on.

Please note that you need to import `bufio` and `os` packages as shown below:

```go
import "bufio"
import "os"
```

Alternatively, since you already have one import (`fmt`), you can add it as a list:

```go
import (
    "bufio"
    "fmt"
    "os"
)
```

The `os.Open()` function returns a pair of values: a file and an error. Returning multiple values from a function is a common pattern in Go, specially for returning errors.

```go
f, err := os.Open(file)
```

The `:=` operator is an assignment operator, but with the special property that it automatically infers the type of the variable(s) based on the value(s) on the right hand side.

Keep in mind that Go is a strongly typed language, but that nice feature saves us the trouble of specifying the type when it's possible to infer it.

In the case above, Go automatically infers the type for both `f` and `err` variables.

When a function returns an error it is a common pattern to check the error immediately afterwards:

```go
    f, err := os.Open(file)
    if err != nil {
        // do something with err
        log.Print("...")
        return
    }
```

Note: It is a good practice to keep the "happy path" aligned to the left, and the sad path to the right (i.e., terminating the function early).

`nil` in Go means no value is assigned to a variable.

The `if` statement executes a statement if the condition is true. It can optionally have an initialization clause just like the `for` statement, and an `else` clause that runs if the condition is false. Please keep in mind that the scope of the variable created will just be the if statement body. For example:

```go
// optional initialization clause
if foo := rand.Intn(2); foo == 0 {
    fmt.Print(foo) // foo is valid here
} else {
    fmt.Print(foo) // and here
}
// but you can't use foo here!
```

Another interesting aspect of the `loadMaze` code is the use of the `defer` keyword. It basically says to call the function after `defer` at the end of the current function. It is very useful for cleanup purposes and in this case we are using it to close the file we've just opened:

```go
func loadMaze(file) error {
    f, err := os.Open(file)
    // omitted error handling
    defer f.Close() // puts f.Close() in the call stack

    // rest of the code

    return nil
    // f.Close is called implicitly
}
```

The next part of the code just reads the file line by line and appends it to the maze slice:

```go
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        maze = append(maze, line)
    }
```

A scanner is a very convenient way to read a file. `scanner.Scan()` will return true while there is something to be read from the file, and `scanner.Text()` will return the next line of input.

The `append` built-in function is responsible for adding a new element to the `maze` slice.

## Task 02: Printing to the Screen

Once we have the maze file loaded into memory we need to print it to the screen.

One way to do that is to iterate over each entry in the `maze` slice and print it. This can be conveniently done with the `range` operator:

```go
func printScreen() {
    for _, line := range maze {
        fmt.Println(line)
    }
}
```

Please note that we are using the `:=` assignment operator to initialize two values: the underscore (_) and the `line` variable. The underscore is just a placeholder for where the compiler would expect a variable name. Using the underscore means that we are ignoring that value.

In the case of the `range` operator, the first return value is the index of the element, starting from zero. The second return value is the value itself.

If we did not write the underscore character to ignore the first value, the range operator would return just the index (and not the value). For example:

```go
for idx := range maze {
    fmt.Println(idx)
}
```

Since in this case we only care about the content and not the index, we can safely ignore the index by assigning it to the underscore.

## Task 03: Updating the game loop

Now that we have both `loadMaze` and `printScreen` functions, we should update the `main` function to initialize the maze and print it on the game loop. See the code below:

```go
func main() {
    // initialise game

    // load resources
    err := loadMaze("maze01.txt")
    if err != nil {
        log.Println("failed to load maze:", err)
        return
    }

    // game loop
    for {
        // update screen
        printScreen()

        // process input

        // process movement

        // process collisions

        // check game over

        // Temp: break infinite loop
        break

        // repeat
    }
}
```

Like always we are keeping the happy path to the left, so if the `loadMaze` function fails we use `log.Println` to log it and then `return` to terminate the program execution. Since we are using a new package, `log`, please make sure it is added to the import section:

```go
import (
    "bufio"
    "fmt"
    "log"
    "os"
)
```

Some IDEs, like `vscode`, can be configured to do this automatically for you.

Note: one could also use `log.Fatalln` for the same effect, but we need to make sure that any deferred calls are executed before exiting the `main` function, and functions in the `log.Fatal` family skip deferred function calls by calling `os.Exit(1)` internally. We don't have any deffered calls in the main function yet, but we will add one in the next chapter.

Now that we've finished the game loop modifications we can run the program with `go run` or compile it with `go build` and run it as a standalone program.

```sh
go run main.go
```

You should see the maze printed to the terminal.

[Take me to step 02!](../step02/README.md)
