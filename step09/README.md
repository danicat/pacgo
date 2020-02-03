# Step 09: Buffer "The String Concatenation Slayer"

In this lesson you will learn how to:

- Concatenate strings using a Buffer

## Overview

In this lesson we will be adding support for multiple lives to the application. We will update the collision tracking code to decrement the number of lives instead of setting lives to 0 on a collision. We will also keep track of the starting player position to respawn player there if they die. Finally, we will add Player emojis to the game scoreboard to track the number of lives remaining instead of displaying lives as an integer value.

## Task 01: Create Point type and update Player struct to use Point type.

We need to track the initial position of the player so we can reset the position after collision with a ghost. We will do this by updating the sprite struct to include `startRow` and `startCol` properties.

```go
type sprite struct {
    row      int
    col      int
    startRow int
    startCol int
}
```

We can then to fill those properties for our player (and ghosts) in the `loadMaze` function:

```go
func loadMaze() error {
    //...omitted for brevity

    for row, line := range maze {
        for col, char := range line {
            switch char {
            case 'P':
                player = sprite{row, col, row, col}
            case 'G':
                ghosts = append(ghosts, &sprite{row, col, row, col})
            case '.':
                numDots++
            }
        }
    }

    return nil
}
```

Note that since we have the additional `startRow` and `startCol` properties we can't do a simple comparisson for collision detection anymore, as the player never starts in the same position as a ghost. We will make sure to reflect that difference in the next session.

## Task 02: Update initial `lives` to be greater than 1 and decrement lives on ghost collision

As a starting point we will set our initial number of lives to 3

```go
var lives = 3
```

We will then update the code that processes collisions to decrement the number of lives by 1 every time a collision occurs. Finally we will check to make sure that we are not out of lives and reset our player emoji to the initial position to restart play.

```go
    // process collisions
    for _, g := range ghosts {
        if player.row == g.row && player.col == g.col {
            lives = lives - 1
            if lives != 0 {
                moveCursor(player.row, player.col)
                fmt.Print(cfg.Death)
                moveCursor(len(maze)+2, 0)
                time.Sleep(1000 * time.Millisecond) //dramatic pause before resetting player position
                player.row, player.col = player.startRow, player.startCol
            }
        }
    }
```

## Task 03: Update scoreboard to display Player emojis corresponding to number of lives

Previously the number of lives was being displayed as an integer in the game scoreboard. We will now be updating the scoreboard to display the number of lives with player emojis. We will be adding a `getLivesAsEmoji` function to concatenate the correct number of player emojis based on the lives remaining in the game. This function creates a buffer and then writes the player emoji string to the buffer based on the number of lives and then returns that value as a string. This function is called in the last line of the `printScreen` function to update the scoreboard.

```go
func printScreen() {
    //...omitted for brevity

    moveCursor(len(maze)+1, 0)

    livesRemaining := strconv.Itoa(lives) //converts lives int to a string
    if cfg.UseEmoji {
        livesRemaining = getLivesAsEmoji()
    }

    fmt.Println("Score:", score, "\tLives:", livesRemaining)
}

//concatenate the correct number of player emojis based on lives
func getLivesAsEmoji() string{
    buf := bytes.Buffer{}
    for i := lives; i > 0; i-- {
        buf.WriteString(cfg.Player)
    }
    return buf.String()
}
```

So why use a buffer? Turns out there are other ways to concatenate strings in Go. The simplest option would be to just use `+` operator to concatenate two strings:

```go
string1 := "pac"
string2 := "go"
pacgo := string1 + string2 //"pacgo"
```

For comparison, this is what the `getLivesAsEmoji` function would look like if we used the `+` operator approach.

```go
func getLivesAsEmoji() string {
    emojiString := ""
    for i := lives; i > 0; i-- {
        emojiString = emojiString + cfg.Player
    }
    return emojiString
}
```

This version of `getLivesAsEmoji` will be less efficient than the version of the function that uses a buffer. Part of the reason for this performance difference is due to memory allocation required for the string concatenation, as we've seen before that strings in Go are immutable.

In the version of the function using the `+` operator, there is a memory allocation operation happening for every iteration of the for loop. While for the buffer version of the function there is only a single memory allocation happening when buffer is initialized. A more detailed example of this performance difference is discussed [here](https://billglover.me/2019/03/13/learn-go-by-concatenating-strings/)


[Take me to step 10!](../step10/README.md)
