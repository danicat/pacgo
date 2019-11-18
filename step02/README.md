# Step 02: Handling Player Input

In this lesson you will learn how to:

- Work with different terminal modes
- Calling external commands from Go code
- Send escape sequences to the terminal
- Read from standard input
- Create a function that returns multiple values

## Overview

In the last step we've learned how to print something to the standard output. Now it's time to learn how to read from standard input.

In this game we will be processing a restricted set of movements: up, down, left and right. Besides those, the only other key we will be using is the escape key, in order to enable the player to exit the game gracefully. The movements will be mapped to the arrow keys.

This step will handle the Esc key and we will see how to process the arrow keys in step 03.

But, before getting into the implementation, we need to know a bit about terminal modes.

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

To enable the cbreak mode we are going to call an external command that controls terminal behaviour, the `stty` command. We are also going to disable terminal echo so we don't polute the screen with the output of key presses.

Here is the definition of our `init`:

```go
func initialise() {
    cbTerm := exec.Command("stty", "cbreak", "-echo")
    cbTerm.Stdin = os.Stdin

    err := cbTerm.Run()
    if err != nil {
        log.Fatalln("unable to activate cbreak mode:", err)
    }
}
```

You will need to add the import "os/exec" if your IDE is not configured to add it automatically.

The `log.Fatalln` function will terminate the program after printing the log, in case of error. This is important here because without the cbreak mode the game is unplayable. As this is the very first function we will call in our program, we are not worried about skipping any deferred calls.

## Task 02: Restoring Cooked Mode

Restoring the cooked mode is a pretty straightforward process. It is the same as enabling the cbreak mode, but with the flags reversed:

```go
func cleanup() {
    cookedTerm := exec.Command("stty", "-cbreak", "echo")
    cookedTerm.Stdin = os.Stdin

    err := cookedTerm.Run()
    if err != nil {
        log.Fatalln("unable to restore cooked mode:", err)
    }
}
```

Now we need to call both functions in the `main` function:

```go
func main() {
    // initialise game
    initialise()
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

The `make` function is a [built-in function](https://golang.org/pkg/builtin/#make) that allocates and initialises objects. It is only used for slices, maps and channels. In this case we are creating an array of bytes with size 100 and returning a slice that points to it.

After the usual error handling (we are just passing the error up on the call stack), we are testing if we read just one byte and if that byte is the escape key. (0x1b is the hexadecimal code that represents Esc).

We return "ESC" if the Esc key was pressed or an empty string otherwise.

Now you may wonder why allocating a buffer of 100 bytes, or why testing the count of exact one byte...

What if the buffer suddenly has 5 elements and one of them is the Esc key? Shouldn't we care to process that? Will that key press be lost?

The short answer is we shouldn't care. Please keep in mind that this is a game. Depending on the processing speed and the length of your keyboard buffer, if we processed events sequentially we could introduce movement lag, i.e., by having a queue of arrow key presses that were not processed yet.

Since we are reading the input on a loop, there is no damage in dropping all the key presses in a queue and just focusing on the last one. That will make the game respond better than if we were concerned about every key press.

## Task 04: Updating the Game Loop

Now it's time to update the game loop to have the `readInput` function called every iteration. Please note that if an error occurs we need to break the loop as well.

```go
// process input
input, err := readInput()
if err != nil {
    log.Print("error reading input:", err)
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

Since we now have a proper game loop, we need to clear the screen after each loop so we have a blank screen for drawing in the next iteration. In order to do that, we are going to use some special "escape sequences".

[Escape sequences](https://en.wikipedia.org/wiki/ANSI_escape_code#Escape_sequences) are called like that because they start with the ESC character (0x1b) followed by one or more characters. Those characters work as commands for the terminal emulator.

You actually don't need to worry about the sequences we are going to use, as we are going to import another package called `simpleansi` that does the work for us:

```go
import "github.com/danicat/simpleansi"
```

---
### A note on external packages

This time we are not importing a package from the standard library, but an external package instead. If you look at `simpleansi`'s [implementation](https://github.com/danicat/simpleansi), you will notice that every function starts with a capital letter, like `ClearScreen` or `MoveCursor`.

That is important in Go because the capitalisation of a word defines if that function or variable has **public** or **private** escope.

Words starting with a lower case character are private to the package defining it, and words starting with an upper case character are public. That may be confusing to people coming from other languages like java, but if you follow naming conventions like "classes (structs) always start with a capital letter" you may end up inadvertedly making every type in your code public, which is probably not what you want.

---

We will update the printScreen function to call `simpleansi.ClearScreen` before printing, so we are sure to be using a blank screen each frame:

```go
func printScreen() {
    simpleansi.ClearScreen()
    for _, line := range maze {
        fmt.Println(line)
    }
}
```

Now run the game again and try hitting the `ESC` key.

Please note that if you hit Ctrl+C by any chance the program will terminate without calling the cleanup function, so you won't be able to see what you are typing in the terminal (because of the `-echo` flag).

If you get into that situation either close the terminal and reopen it or just run the game again and exit gracefully using the `ESC` key.

[Take me to step 03!](../step03/README.md)