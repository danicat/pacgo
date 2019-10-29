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
func loadConfig() error {
    f, err := os.Open("config.json")
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
err = loadConfig()
if err != nil {
    log.Println("Error loading configuration:", err)
    return
}
```

## Task 02: Adjusting the horizontal displacement

We need to adapt the `moveCursor` function to correct the horizontal displacement when the emoji flag is set:

```go
func moveCursor(row, col int) {
    if cfg.UseEmoji {
        fmt.Printf("\x1b[%d;%df", row+1, col*2+1)
    } else {
        fmt.Printf("\x1b[%d;%df", row+1, col+1)
    }
}
```

Scaling the `col` value by 2 times will ensure we position every character in the right place. It will also have the side effect of making the maze look bigger.

## Task 03: Replace hardcoded characters with configuration

The final part is to replace the hardcoded characters with their config counterparts in the `printScreen` function:

```go
func printScreen() {
    clearScreen()
    for _, line := range maze {
        for _, chr := range line {
            switch chr {
            case '#':
                fmt.Print(cfg.Wall)
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

As an added bonus, let's add a game over sprite within the game over condition. Note that this will work only if your `printScreen` call is at the beginning of the game loop before anything else is processed:

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

We have emojis! How great is that? :)

[Take me to step 08!](../step08/README.md)
