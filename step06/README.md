# Step 06: Making things real(time)

In this lesson you will learn how to:

- Use goroutines
- Use anonymous functions (lambdas)
- Use channels
- Use the select statement to read channels async
- Use package time

## Overview

At this point of the tutorial we kind of have a complete game: it has a clear objective, winning and losing conditions and the player input works correctly.

But it has one major issue: the enemies move only when the player moves. That doesn't look like very gamey to me, so let's do this properly.

This issue happens because the read input is a blocking operation. We need to make it asynchronous somehow... if only we had some functionality to run things async in go... Oh, wait! We do! :)

Here comes the fabulous channels and goroutines to the rescue!

Goroutines are similar to threads, but they are much more lightweight. Under the hood, the go runtime spawns threads that will handle the goroutines, but a single thread can manage several goroutines, so their relation is bigger than 1:1.

But that's not the best part. The go language design makes it very easy to spawn a goroutine: you just need to add the keyword `go` before the function call and the function will run on a separate goroutine in an asynchronous manner.

Have a look at the code below:

```go
func main() {
    go fmt.Println("hello")
    go fmt.Println("world")
}
```

This code has three goroutines: the first one is the one that runs the `main` function, the second one is the one that prints `hello` and the third one is the one that prints `world`.

One important thing about goroutines is that since we are doing things async, it's safe to assume that the previous program will produce no output. That's because the `main` function has a high probability of terminating the program before any of the two goroutines are executed (because we have some overhead to launch the goroutines).

We could introduce a delay to the main function:

```go
func main() {
    go fmt.Println("hello")
    go fmt.Println("world")
    time.Sleep(100 * time.Millisecond)
}
```

That would guarantee that the goroutines would run, as we expect them to be faster than 100ms, but still, the output of this program is unpredictable, as we cannot count on the order that the goroutines are executed. 

Once a `go` statement is executed, the responsibility for scheduling the goroutine for execution is passed to the go runtime. We don't have control over this and we can never assume a specific order of execution. Keep that in mind when writing async code.

In addition to goroutines, we also have the channel constructs. Channels allow us to communicate with goroutines by passing or receiving values. Or both.

To create a channel, we use the `make` built-in function:

```go
ch := make(chan int)
```

Each channel has a type, and optionally a buffer size. If no size is specified it is assumed to be 1.

Reading and writing to a channel can be a blocking operation, unless the channel is empty.

You write to a channel using the arrow operator:

```go
// something is written to ch
ch <- something
```

In the scenario above, if ch is empty the operation won't block, but if it's full it will block until the channel is consumed on the other side.

Similarly, reading from a channel also uses the arrow operator:

```go
// on a different goroutine
foo := <-ch
```

When designing async processing we must be careful that two goroutines don't depend on each other in a way that they can both be held in a blocking state or produce inconsistent results. To know more about deadlocks and race conditions, please see [this answer](https://stackoverflow.com/a/3130212/4893628) on StackOverflow.

## Task 01: Refactoring the input code

Now that we know the basics about goroutines and channels, let's see them in action. First, let's remove the input handling code from the game loop and insert the code below before the start of the loop.

```go
func main() {
    // init code omitted for brevity...

    // process input (async)
    input := make(chan string)
    go func(ch chan<- string) {
        for {
            input, err := readInput()
            if err != nil {
                log.Println("error reading input:", err)
                ch <- "ESC"
            }
            ch <- input
        }
    }(input)

    // game loop
    for {
        // loop code...
    }
}
```

This code will create a channel called `input` and pass it as a parameter to an anonymous function that is invoked with the `go` statement. That's a very common pattern for async processing in go.

The anonymous function then creates an infinite loop where it waits for input and writes it to the channel `ch` (given by the function parameter). In case of error, we just return the "ESC" code as we know this will terminate the program.

In the game loop we will replace the code that processes the player movement with the code below:

```go
// process movement
select {
case inp := <-input:
    if inp == "ESC" {
        lives = 0
    }
    movePlayer(inp)
default:
}
```

Imagine that the select statement is just like a switch statement, but for channels. This select statement has a non-blocking nature, because it has a default clause. This means that if the `input` channel has something to be read it will be read, otherwise the `default` case is processed, which in this case is an empty block.

Finally, since we've moved the "ESC" logic to the block above, we will remove it from the game over conditions (as the `lives <= 0` already satisfies it).

We will also introduce a delay of 200ms. Since now we are not waiting for input anymore the game will run too fast without it. The relevant snippet is below:

```go
    // update screen
    printScreen()

    // check game over
    if numDots == 0 || lives <= 0 {
        break
    }

    // repeat
    time.Sleep(200 * time.Millisecond)
```

Try running the game now. Much more exciting, isn't it? :)

[Take me to step 07!](../step07/README.md)
