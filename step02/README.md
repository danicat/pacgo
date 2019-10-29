# Step 02: Handling Player Input

In this lesson you will learn how to:

- Work with different terminal modes
- Send escape sequences to the terminal
- Read from standard input
- Create a function that return multiple values

## Overview

In the last step we've learned how to print something to the standard output. Now it's time to learn how to read from standard input. 

In this game we will be processing a restricted set of movements: up, down, left and right. Besides those, the only other key we will be needing is the game end key, in order to enable the player to quit the game gracefully.

We will map the escape key (Esc) on the keyboard to do this graceful termination. The movements will be mapped to the arrow keys.

This step will handle the Esc key and we will see how to process the arrow keys in step 03.

But before getting into the implementation we need to know a bit about terminal modes.

## Intro to terminal modes

Terminals can run in three possible [modes](https://en.wikipedia.org/wiki/Terminal_mode): 

1. Cooked Mode
2. Cbreak Mode
3. Raw Mode

The cooked mode is the one that we are used to use. In this mode every input that the terminal receives is preprocessed, meaning that the system intercepts special characters to give them special meaning.

Note: Special characters include backspace, delete, Ctrl+D, Ctrl+C, arrow keys and so on...

The raw mode is the opposite: data is passed as is, without any kind of preprocessing.

The cbreak mode is the middle ground. Some characters are preprocessed and some are not. For instance, Ctrl+C still results in program abortion, but the arrow keys are passed to the program as is.

We will use the cbreak mode to allow us to handle the escape sequences corresponding to the escape and arrow keys.

## Task 01: Enabling Cbreak Mode

To enable the cbreak mode we are going to take advantage of an `init` function.

We said previously that the `main` function is the entrypoint of a given program. Besides that, we can also have an `init` function that perform initialization steps before the runtime calls the `main` function.

Additionally, we can have an `init` function for library packages to perform any necessary initialization. It's useful to note that you shouldn't count on the order of which the `init` functions are called in the scenario where you have multiple packages with multiple `init` functions.

Here is the definition of our `init`:

```go
func init() {
    cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
    cbTerm.Stdin = os.Stdin

    err := cbTerm.Run()
    if err != nil {
        log.Fatalln("Unable to activate cbreak mode terminal:", err)
    }
}
```

In the function above we are actually invoking another program that modifies the terminal configuration, called `stty`. We are both enabling the cbreak mode and disabling the cursor echo.

The `log.Fatalf` function will terminate the program after printing the log, in case of error. This is important here because without the cbreak mode the game is unplayable.

## Task 02: Restoring Cooked Mode

Restoring the cooked mode is a pretty straightfoward process. It is the same as enabling the cbreak mode, but with the flags reversed:

```go
func cleanup() {
    cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
    cookedTerm.Stdin = os.Stdin

    err := cookedTerm.Run()
    if err != nil {
        log.Fatalln("Unable to activate cooked mode terminal:", err)
    }
}
```

Note that this `cleanup` function doesn't have a special meaning in Go like the `init` one, so we must explicitly call it in our `main` function. We can either call it at the end of the function or use the `defer` statement as shown below:

```go
func main() {
    // initialize game
    defer cleanup()

    // load resources
    // ...
```

## Task 03: Reading from Stdin

The process of reading from the standard input involves calling the function `os.Stdin.Read` with a given read buffer.

The `os.Stdin.Read` returns two values: the number of bytes read and an error value. Have a look at the code for the `readInput` function below:

```go
func readInput() (string, error) {
    buffer := make([]byte, 100)

    cnt, err := os.Stdin.Read(buffer)
    if err != nil {
        return "", err
    }

    if cnt == 1 && buffer[0] == 0x1b {
        return "ESC", nil
    }

    return "", nil
}
```

The `make` function is a [built-in function](https://golang.org/pkg/builtin/#make) that allocates and initializes objects. It is only used for slices, maps and channels. In this case we are creating an array of bytes with size 100 and returning a slice that points to it.

After the usual error handling (we are just passing the error up on the call stack), we are testing if we read just one byte and if that byte is the escape key. (0x1b is the hexadecimal code that represents Esc).

We return "ESC" if the Esc key was pressed or an empty string otherwise.

Now you may wonder why allocating a buffer of 100 bytes, or why testing the count of exact one byte... 

What if the buffer suddenly has 5 elements and one of them is the Esc key? Shouldn't we care to process that? Will that key press be lost?

The short answer is we shouldn't care. Please keep in mind that this is a game. Depending on the processing speed and the length of your keyboard buffer, if we processed events sequentially we could introduce movement lag, ie, by having a queue of arrow key presses that were not processed yet.

Since we are reading the input on a loop, there is no damage in dropping all the key presses in a queue and just focusing on the last one. That will make the game response work better than if we were concerned about every key press.

## Task 04: Updating the Game Loop

Now it's time to update the game loop to have the `readInput` function called every iteration. Please note that on the advent of an error we need to break the loop as well.

```go
// process input
input, err := readInput()
if err != nil {
    log.Printf("Error reading input: %v", err)
    break
}
```

Finally, we can get rid of that permanent `break` statement and start testing for the "ESC" key press.

```go
if input == "ESC" {
    break
}
```

## Task 05: Clearing the Screen

Since we now have a proper game loop, we need to clear the screen after each loop so we have a blank screen for drawing in the next iteration. In order to do that, we are going to use some special escape sequences.

```go
func clearScreen() {
    fmt.Print("\x1b[2J")
    moveCursor(0, 0)
}

func moveCursor(row, col int) {
    fmt.Printf("\x1b[%d;%df", row+1, col+1)
}
```

[Escape sequences](https://en.wikipedia.org/wiki/ANSI_escape_code#Escape_sequences) are called like that because they start with the ESC character (0x1b) followed by one or more characters. Those characters work as commands for the terminal emulator.

We are using above two commands: one for clearing the screen and other for moving the cursor to a given position.

We will update the printScreen function to call clearScreen before printing, so we are sure to be using a blank screen each frame:

```go
func printScreen() {
    clearScreen()
    for _, line := range maze {
        fmt.Println(line)
    }
}
```

Now run the game again and try hiting the `ESC` key.

Please note that if you hit Ctrl+C by any chance the program will terminate without calling the cleanup function, so you won't be able to see what you are typing in the terminal (because of the `-echo` flag).

If you get into that situation either close the terminal and reopen it or just run the game again and exit gracefully using the `ESC` key.

[Take me to step 03!](../step03/README.md)