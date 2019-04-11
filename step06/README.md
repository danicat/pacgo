# Step 06: Making things real(time)

In this lesson you will learn how to:

- Use gorotines
- Use channels
- Use the select statement to read channels async
- Use package time

## Overview

At this point of the tutorial we kind of have a complete game: it has a clear objective, winning and losing conditions and process the player input correctly.

But it has one major issue: the enemies move only when the player moves. That doesn't look like very gamey to me, so let's do this properly.

This issue happens because the read input is a blocking operation. We need to make it assynchronous somehow... if only we have some functionally to run things async... luckly we do have!

Here comes the fabulous channels and goroutines to the rescue!

## Task 01: Refactoring the input code

```go
func main() {
    // init code omitted for brevity...

    // process input (async)
    input := make(chan string)
    go func(ch chan<- string) {
        for {
            input, err := readInput()
            if err != nil {
                log.Printf("Error reading input: %v", err)
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

## Task 02: Introducing a delay