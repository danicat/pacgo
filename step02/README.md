# Step 02: Handling Player Input

In this lesson you will learn how to:

- Read from standard input
- Create a function that return multiple values

## Overview

In the last step we've learned how to print something to the standard output. Now it's time to learn how to read from standard input. 

In this game we will be processing a restricted set of movements: up, down, left and right. Besides those, the only other key we will be needing is the game end key, in order to enable the player to quit the game gracefully.

We will map the escape key (Esc) on the keyboard to do this graceful termination. The movements will be mapped to the arrow keys.

This step will handle the Esc key and we will see how to process the arrow keys in step 04.

## Task 01: Reading from Stdin

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

The `make` function is a [built-in function](https://golang.org/pkg/builtin/#make) that allocates and initializes objects. It is only used for slices, maps and channels. In this case we are creating an array of bytes with size 100 and returning a slice that pointers to it.

After the usual error handling (we are just passing the error up on the call stack), we are testing if we read just one byte and if that byte is the escape key. (0x1b is the hexadecimal code that represents Esc).

We return "ESC" if the Esc key was pressed or an empty string otherwise.

Now you may wonder why allocating a buffer of 100 bytes, or why testing the count of exact one byte... 

What if the buffer suddenly has 5 elements and one of them is the Esc key? Shouldn't we care to process that? Will that key press be lost?

The short answer is we shouldn't care. Please keep in mind that this is a game. Depending on the processing speed and the lenght of your keyboard buffer, if we processed events sequentially we could introduce movement lag, ie, by having a queue of arrow key presses that were not processed yet.

Since we are reading the input on a loop, there is no damage in droping all the key presses in a queue and just focusing on the last one. That will make the game response work better than if we were concerned about every key press.

## Task 02: Updating the Game Loop

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