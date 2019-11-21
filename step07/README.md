# Step 07: Finally, emojis!

In this lesson you will learn how to:

- Load a json file
- Print emojis!!!

## Overview

So, we've managed to create a proper game in the terminal. But I've promised emojis, where are they? Well, the time has come, finally!

In this step we will create a file called `config.json`. In this file we will store mappings for each symbol we use in our game. In 2D games, we usually call the moving pieces "sprites".

Since most of the terminals nowadays support unicode, we can use emojis as our sprites without needing to resort to any graphical library.

The provided `config.json` file should look like this:

```json
{
    "player": "😋",
    "ghost": "👻",
    "wall": "🌵",
    "dot": "🧀",
    "pill": "🍹",
    "death": "💀",
    "space": "  ",
    "use_emoji": true
}
```

This is the default mapping but please feel free to toy with the entire emoji palette. We have infinite possibilities!

One important aspect about the config file is the `use_emoji` configuration. We are using this flag to signal to the game when we are using emojis. This is necessary because emojis generally use more than one character in the screen (most of them use 2).

Using that flag we can have alternate code paths that make adjustments to compensate that difference. Otherwise the maze would look distorted.

## Task 01: Load a json

Go has support for loading json in the standard library.

We first need to define a struct to hold the json data. The text between the backticks (\`) is called a `struct tag`. It is used by the json decoder to know which field of the struct corresponds to each field in the json file.

```go
// Config holds the emoji configuration
type Config struct {
    Player   string `json:"player"`
    Ghost    string `json:"ghost"`
    Wall     string `json:"wall"`
    Dot      string `json:"dot"`
    Pill     string `json:"pill"`
    Death    string `json:"death"`
    Space    string `json:"space"`
    UseEmoji bool   `json:"use_emoji"`
}

var cfg Config
```

Note that we used public members for the `Config` struct. That is required for the json decoder to work.

The code below parses the json and stores it in the `cfg` global variable.

```go
func loadConfig(file string) error {
    f, err := os.Open(file)
    if err != nil {
        return err
    }
    defer f.Close()

    decoder := json.NewDecoder(f)
    err = decoder.Decode(&cfg)
    if err != nil {
        return err
    }

    return nil
}
```

Now add the `loadConfig` call in the initialization part of the main function, after `loadMaze`:

```go
err = loadConfig("config.json")
if err != nil {
    log.Println("failed to load configuration:", err)
    return
}
```

## Task 02: Adjusting the horizontal displacement

We need to create a custom `moveCursor` function to correct the horizontal displacement when the emoji flag is set:

```go
func moveCursor(row, col int) {
    if cfg.UseEmoji {
        simpleansi.MoveCursor(row, col*2)
    } else {
        simpleansi.MoveCursor(row, col)
    }
}
```

Make sure you replace all calls to `simpleansi.MoveCursor` with calls to the new `moveCursor` function (except for the ones inside the new function).

Scaling the `col` value by 2 times will ensure we position every character in the right place. It will also have the side effect of making the maze look bigger.

## Task 03: Replace hardcoded characters with configuration

The final part is to replace the hardcoded characters with their config counterparts in the `printScreen` function. We are also going to use the `simpleansi.WithBlueBackground` function to change the colour of the walls to make it more representative of the original game.

```go
func printScreen() {
    simpleansi.ClearScreen()
    for _, line := range maze {
        for _, chr := range line {
            switch chr {
            case '#':
                fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
            case '.':
                fmt.Print(cfg.Dot)
            default:
                fmt.Print(cfg.Space)
            }
        }
        fmt.Println()
    }

    moveCursor(player.row, player.col)
    fmt.Print(cfg.Player)

    for _, g := range ghosts {
        moveCursor(g.row, g.col)
        fmt.Print(cfg.Ghost)
    }

    moveCursor(len(maze)+1, 0)
    fmt.Println("Score:", score, "\tLives:", lives)
}
```

## Task 04: Game over and pills

As an added bonus, let's add a game over sprite within the game over condition.

```go
// check game over
if numDots == 0 || lives == 0 {
    if lives == 0 {
        moveCursor(player.row, player.col)
        fmt.Print(cfg.Death)
        moveCursor(len(maze)+2, 0)
    }
    break
}
```

Also, let's add the code to treat the power up pill as a dot that worth more points, as a placeholder for the actual power up mechanics. We are just doing this now to have a sense of a complete game, but we'll come to implement the proper power up mechanics at a later step.

```go
func movePlayer(dir string) {
    player.row, player.col = makeMove(player.row, player.col, dir)

    removeDot := func(row, col int) {
        maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
    }

    switch maze[player.row][player.col] {
    case '.':
        numDots--
        score++
        removeDot(player.row, player.col)
    case 'X':
        score += 10
        removeDot(player.row, player.col)
    }
}
```

Once interesting thing about the code above is that we are defining an inline function to do the removal of both the dot and the X from the game when we have a collision. We could also have repeated the code, but this makes it more readable and maintainable.

We have emojis! How great is that? :)

[Take me to step 08!](../step08/README.md)
