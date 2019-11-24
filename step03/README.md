# Step 03: Adding Movement

In this lesson you will learn how to:

- Create a struct
- Use the switch statement
- Handle the arrow keys
- Use named return values

## Overview

We have a maze, we can quit the game gracefully... but nothing very exciting is happening, right? So let's spice this thing up a bit and add some movement!

In this step we are adding the player character and enabling its movement with the arrow keys.

## Task 01: Tracking player position

The first step in our journey is to create a variable to hold the player data. Since we will be tracking 2D coordinates (row and column), we will define a struct to hold that information:

```go
type sprite struct {
    row int
    col int
}

var player sprite
```

We are also defining the player as a global variable, just for the sake of simplicity.

Next we need to capture the player position as soon as we load the maze, in the `loadMaze` function:

```go
// traverse each character of the maze and create a new player when it locates a `P`
for row, line := range maze {
    for col, char := range line {
        switch char {
        case 'P':
            player = sprite{row, col}
        }
    }
}
```

Note that this time we are using the full form of the `range` operator, as we are interested in which row and column we found the player.

Here is the complete `loadMaze` just for reference:

```go
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

    for row, line := range maze {
        for col, char := range line {
            switch char {
            case 'P':
                player = sprite{row, col}
            }
        }
    }

    return nil
}
```

---

### Optional: A note about visibility

We are keeping things simple here just for the sake of the tutorial. Since everything is a single file we are not taking into account the visibility of variables, i.e., if they are public or private.

Nevertheless, Go has an interesting mechanic in regards to defining visibility. Instead of having a public keyword, it considers public every symbol whose name starts with a capital letter. On the other hand, if a name starts with a lowercase character, it is a private symbol.

That's why every library function name we've used so far begins with a capital letter. That's also why your IDE may complain about missing comments if you define any variable, function or type with an initial uppercase character. In the Go idiom, public symbols should always be commented, as those are later extracted to become the package documentation.

In this particular case, we are using lowercase symbols for all our variables, types and functions since it doesn't make any sense to export a symbol from the package `main`.

---

## Task 02: Handling arrow key presses

Next, we need to modify `readInput` to handle the arrow keys:

```go
if cnt == 1 && buffer[0] == 0x1b {
    return "ESC", nil
} else if cnt >= 3 {
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
```

The escape sequence for the arrow keys are 3 bytes long, starting with `ESC+[` and then a letter from A to D.

We now need a function to handle the movement:

```go
func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
    newRow, newCol = oldRow, oldCol

    switch dir {
    case "UP":
        newRow = newRow - 1
        if newRow < 0 {
            newRow = len(maze) - 1
        }
    case "DOWN":
        newRow = newRow + 1
        if newRow == len(maze) {
            newRow = 0
        }
    case "RIGHT":
        newCol = newCol + 1
        if newCol == len(maze[0]) {
            newCol = 0
        }
    case "LEFT":
        newCol = newCol - 1
        if newCol < 0 {
            newCol = len(maze[0]) - 1
        }
    }

    if maze[newRow][newCol] == '#' {
        newRow = oldRow
        newCol = oldCol
    }

    return
}
```

Note: If you are used to the switch statement in other languages, please beware that in Go there is an implicit `break` after each `case` condition. So we don't need to explicitly break after each block. If we want to fall through the next `case` block we can use the `fallthrough` keyword.

The function above takes advantage of `named return values` to return the new position (`newRow` and `newCol`) after the move. Basically the function "tries"  the move first, and if by any chance the new position hits a wall (`#`) the move is cancelled.

It also handles the property that if the character moves outside the range of the maze it appears on the opposite side.

The last piece in the movement puzzle is to define a function to move the player:

```go
func movePlayer(dir string) {
    player.row, player.col = makeMove(player.row, player.col, dir)
}
```

## Task 03: Updating the maze

We have all the movement logic in place, but we need to make the screen reflect that. We will refactor the `printScreen` function to print only the things that we want to print, instead of the whole map.

That will give us more control, enabling us to print the player at an arbitrary position with the `moveCursor` function. See the code below:

```go
func printScreen() {
    simpleansi.ClearScreen()
    for _, line := range maze {
        for _, chr := range line {
            switch chr {
            case '#':
                fmt.Printf("%c", chr)
            default:
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }

    simpleansi.MoveCursor(player.row, player.col)
    fmt.Print("P")


    // Move cursor outside of maze drawing area
    simpleansi.MoveCursor(len(maze)+1, 0)
}
```

For the time being, we are ignoring anything that is not a wall or the player.

## Task 04: Animation!

Finally, we need to call `movePlayer` from the game loop:

```go
// game loop
for {
    // update screen
    printScreen()

    // process input
    input, err := readInput()
    if err != nil {
        log.Println("error reading input:", err)
        break
    }

    // process movement
    movePlayer(input)

    // process collisions

    // check game over
    if input == "ESC" {
        break
    }

    // repeat
}
```

We are good to Go!

[Take me to step 04!](../step04/README.md)
