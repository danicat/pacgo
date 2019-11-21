# Step 04: Ghosts!

In this lesson you will learn how to:

- Create a map (dictionary)
- Generate random numbers
- Use pointers

## Overview

Now that we can move our player, it's time to do something about our enemies (ghosts).

We will use the same move mechanic as the player one, the `makeMove` function, but instead of reading input from the keyboard we will use a simple algorithm: generate a random number between 0 and 3 and assign a direction to each of those values.

If the Ghost hits a wall it doesn't matter, it will just try again on the next iteration.

## Task 01: Making Ghosts

Just like we've created a struct to hold our player data, we will create a similar one for ghosts. The only difference is that instead of holding a player global variable in memory we will have a slice of pointers to sprites. That way we can update each ghost's position in a very efficient way.

First, the declaration:

```go
var ghosts []*sprite
```

Note the `*` symbol denoting that `[]*sprite` is a slice of **pointers** to sprite objects.

Next, loading. In the `loadMaze` function, add a new case to the switch statement for handling `G` symbols on the map:

```go
for row, line := range maze {
    for col, char := range line {
        switch char {
        case 'P':
            player = sprite{row, col}
        case 'G':
            ghosts = append(ghosts, &sprite{row, col})
        }
    }
}
```

Please note the ampersand (`&`) operator. This means that instead of adding a sprite object to the slice we are adding a pointer to it.

Go is a garbage collected language, which means it can automatically de-allocate a piece of memory when it is no longer used. Because of that we can use pointers in a much safer way than, for instance, in C++. We are also not allowed to do math on pointers. In effect, a pointer in Go works almost like a reference.

Now, since we are handling `G`s in the `loadMaze` function we also need to print them in `printScreen`. Just add the following block after printing the player:

```go
for _, g := range ghosts {
    simpleansi.MoveCursor(g.row, g.col)
    fmt.Print("G")
}
```

## Task 02: A very smart AI

We've mentioned before that we would be using a random number generator to control our ghosts. That sounds way more complex than it actually is. Have a look at the code:

```go
func drawDirection() string {
    dir := rand.Intn(4)
    move := map[int]string{
        0: "UP",
        1: "DOWN",
        2: "RIGHT",
        3: "LEFT",
    }
    return move[dir]
}
```

The function `rand.Intn` from the `math/rand` package generates a random number between the interval `[0, n)`, where `n` is the parameter given to the function. (Note that the interval is open ended, so `n` is not included).

We are using a trick to map the integer numbers to the actual movements using a `map`. A map is a data structure that maps one value to another. I.e., in the case above, the map `move` maps an integer to a string.

## Task 03: Let's add some movement!

Finally, we need a function to process the ghost movement. The `moveGhosts` function is presented below:

```go
func moveGhosts() {
    for _, g := range ghosts {
        dir := drawDirection()
        g.row, g.col = makeMove(g.row, g.col, dir)
    }
}
```

Now update the game loop to call `moveGhosts`:

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
    moveGhosts()

    // process collisions

    // check game over
    if input == "ESC" {
        break
    }

    // repeat
}
```

We are done! Now we have ghosts that move! How scary -_-'''

[Take me to step 05!](../step05/README.md)
