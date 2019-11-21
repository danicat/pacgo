# Step 08: Command line parameters

In this lesson you will learn how to:

- Add flags to a command line application

## Overview

In the last lesson we added a configuration file `config.json` to handle our emoji translation. We also have a file named `config_noemoji.json` that translates to the original representation of the game.

We also use the `maze01.txt` file for our maze representation. All those names are written directly to the source code, but dealing with these files in a hard coded way is not ideal, so we will change that.

## Task 01: Create flags for each file

The `flag` package of the standard library is the one responsible for handling command line flags. We are going to use it to create two flags: `--config-file` and `--maze-file`.

At the beginning of the file, just after the imports, add the following global variables.

```go
var (
    configFile = flag.String("config-file", "config.json", "path to custom configuration file")
    mazeFile   = flag.String("maze-file", "maze01.txt", "path to a custom maze file")
)
```

The `String` function of the `flag` package accepts three parameters: a flag name, a default value and a description (to be exhibited when `--help` is used). It returns a pointer to a string which will hold the value of the flag.

Please note that this value is only filled after calling `flag.Parse`, which should be called from the `main` function:

```go
func main() {
    flag.Parse()

    // initialise game
    initialise()
    defer cleanup()

    // rest of the function omitted...
}
```

Please note that we are calling `flag.Parse()` as the very first thing in the program. We want to do that because we want the flags to be parsed **before** changing the console to `cbreak` mode.

When the flag is parsed in case of error it calls `os.Exit`, which means our `cleanup` function wouldn't be called leaving the terminal without echo and still in cbreak mode, which can be quite inconvenient.

With this change, by controlling the order things are called, we are making sure we init the cbreak mode only when the flags are parsed successfully.

## Task 02: Replacing the hard coded files with the flags

We've already handled the parsing, now we need to replace the hard coded values with their flag equivalents.

This is done by replacing the hard coded value with the value of the flag (please note the de-reference operator, as the flags are pointers).

In `main`:

```go
    // load resources
    err := loadMaze(*mazeFile)
    if err != nil {
        log.Println("failed to load maze:", err)
        return
    }

    err = loadConfig(*configFile)
    if err != nil {
        log.Println("failed to load configuration:", err)
        return
    }
```

Now try running in the command line:

```sh
go build
./step08 --help
```

You should see something like:

```sh
$ ./step08 --help
Usage of ./step08:
  -config-file string
        path to custom configuration file (default "config.json")
  -maze-file string
        path to a custom maze file (default "maze01.txt")
```

Now try running `step08` with `--config-file config_noemoji.json` first, and `--config-file config.json` later to see the difference. Better with emojis right?

You can also try copying `maze01.txt` to a new file and editing it to experiment.

Maybe you can create your own themes now... try visiting [Full Emoji List](https://unicode.org/emoji/charts/full-emoji-list.html) for inspiration. :)

[Take me to step 09!](../step09/README.md)
