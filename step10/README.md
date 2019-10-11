# Step 08: Command line parameters

In this lesson you will learn how to:

- Concatenate strings using a Buffer

## Overview

In this lesson we will be adding support for multiple lives to the application. We will update the collision tracking code to decrement the number of lives instead of setting lives to 0 on a collision. We will also add a player position reset to restart play after collision with a ghost. Finally, we will add Player emojis to the game scoreboard to track the number of lives remaining instead of displaying lives as an integer value.

## Task 01: Create initialPosition variable

We need to track the initial position of the player so we can reset the position after collision with a ghost. This variable will just hold the row and column of the player at the very start of the game. Since with already have a Player struct with this information we can just declare another variable of Player type to store our initial position data.

```go
var initialPosition Player
```

We can then set our `initialPosition` in the `loadMaze` function

```go
func loadMaze() error {
    //...omitted for brevity

    for row, line := range maze {
        for col, char := range line {
            switch char {
            case 'P':
                player = Player{row, col}
                initialPosition = Player{row, col}
            case 'G':
                ghosts = append(ghosts, &Ghost{row, col})
            case '.':
                numDots++
            }
        }
    }

    return nil
}
```

## Task 02: Update initial `lives` to be greater than 1 and decrement lives on ghost collision

As a starting point we will set our initial number of lives to 3

```go
var lives = 3
```
We will then update the code that processes collisions to decrement the number of lives by 1 everytime a collision occurs. Finally we will check to make sure that we are not out of lives and reset our player emoji to the initial position to restart play. 

```go
    // process collisions
    for _, g := range ghosts {
        if player.row == g.row && player.col == g.col {
            lives = lives - 1
            if lives != 0 {
                moveCursor(player.row, player.col)
                fmt.Printf(cfg.Death)
                moveCursor(len(maze)+2, 0)
                time.Sleep(1000*time.Millisecond) //dramatic pause before reseting player position
                player = initialPosition
            }
        }
    }
```

## Task 03: Update scoreboard to display Player emojis corresponding to number of lives

Previously the number of lives was being displayed as an integer in the game scoreboard. We will now be updating the scoreboard to display the number of lives with player emojis. We will be adding a `drawLives` function to concatenate the correct number of player emojis based on the lives remaining in the game. This function creates a buffer and then writes the player emoji string to the buffer based on the number of lives and then returns that value as a string. This function is called in the last line of the `printScreen` function to update the scoreboard.

```go
func printScreen() {
    //...omitted for brevity

    moveCursor(len(maze)+1, 0)
    fmt.Printf("Score: %v\tLives: %v\n", score, drawLives())
}

//concatenate the correct number of player emojis based on lives
func drawLives() string{
    buf := bytes.Buffer{}
    for i := lives; i > 0; i-- {
        buf.WriteString(cfg.Player)
    }
    return buf.String()
}
```

[Take me to Next Step!](../stepxx/README.md)