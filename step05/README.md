# Step 05: Game over man, game over!

In this lesson you will learn how to:

- Use the fallthrough statement in switch blocks
- Work with slices

## Overview

We are almost there. We have both player movement and ghost movement. But our ghosts are still inoffensive.

It's time to add some danger to this game. Also, with great risks should come great rewards, so we'll also be tackling the game win condition, by clearing the board of all its dots.

## Task 01: Preparation

For the game win condition we need to keep track of how many dots we have on the board, and declare win when this number is zero.

We will need a mechanic to remove the dots from the board once the player stands over them. We will also keep track of a score to show the player.

For the game over scenario, we will be giving the player one life and when a ghost hits them this life is zeroed. We then test for zero lives in the game loop to terminate the game. (It should be pretty straightforward to add support for multiple lives, but we will do this at a later step).

Add the following global variables to track all of the above:

```go
var score int
var numDots int
var lives = 1
```

Next, we need to initialize the `numDots` variable in `loadMaze`. We just need a new case in the switch that handles the `.` character:

```go
for row, line := range maze {
    for col, char := range line {
        switch char {
        case 'P':
            player = sprite{row, col}
        case 'G':
            ghosts = append(ghosts, &sprite{row, col})
        case '.':
            numDots++
        }
    }
}
```

Now we need to update the `printScreen` function to print the dots again. This is an interesting case for the `fallthrough` statement:

```go
func printScreen() {
    simpleansi.ClearScreen()
    for _, line := range maze {
        for _, chr := range line {
            switch chr {
            case '#':
                fallthrough
            case '.':
                fmt.Printf("%c", chr)
            default:
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
    // rest of the function omitted for brevity...
}
```

Finally, at the end of the `printScreen` function let's add our score and lives panel:

```go
func printScreen() {
    // code omitted...

    // print score
    simpleansi.MoveCursor(len(maze)+1, 0)
    fmt.Println("Score:", score, "\tLives:", lives)
}
```

## Task 02: Game over

To process game over is pretty straightforward. At any given moment in time, we are killing the player if they are in the same spot as a ghost. We will add the code that detects this to the game loop. We are also modifying the game quit condition adding `lives <= 0` and `numDots == 0`:

```go
// game loop
for {
    // code ommited...

    // process collisions
    for _, g := range ghosts {
        if player == *g {
            lives = 0
        }
    }

    // check game over
    if input == "ESC" || numDots == 0 || lives <= 0 {
        break
    }

    // repeat
}
```

Please note that the more verbose way of checking the player position is `player.row == g.row && player.col == g.col`, but since both player and ghost are sprites they can use a simple comparison `player == *g`. We still need to dereference `g` because we can't compare pointer and non pointer types.

## Task 03: Game win

We are now just missing the code to remove the dots from the game and increment the score.

We will add this code to the `movePlayer` function:

```go
func movePlayer(dir string) {
    player.row, player.col = makeMove(player.row, player.col, dir)
    switch maze[player.row][player.col] {
    case '.':
        numDots--
        score++
        // Remove dot from the maze
        maze[player.row] = maze[player.row][0:player.col] + " " + maze[player.row][player.col+1:]
    }
}
```

The code above works as follows: first we make the movement. Then we check which character is in the same spot as the player. If it is a dot, we decrement the total number of dots (`numDots`), increment the score and remove the dot from the board.

It's worthwhile to mention that strings in Go are immutable. We couldn't simply assign a space to a given position in a string. It wouldn't work.

Hence, we are using here a trick by creating a new string composed by two slices of the original string. The two slices together make the exact same original string except for one position, that we replace with a space.

For more information about slices, please see [here](https://blog.golang.org/go-slices-usage-and-internals).

Now we have both game win and game over conditions. Try building a map with just a couple of dots and test the game win. Hit a ghost to test the game over. We are making progress!

(Tip: the maze01.txt at the step05 folder has only 3 dots.)

[Take me to step 06!](../step06/README.md)
