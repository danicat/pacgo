# pacgo

A Pac Man clone written in Go (with emojis!)

## Introduction

Welcome to Pac Go! This project is a tutorial to introduce people to the [Go programming language](https://golang.org).

## Why a new tutorial?

We have a lot of great tutorials out there, but the whole idea about this tutorial is to make something different and fun while learning Go. Yes, we need to know about APIs and CRUDs on a daily basis, but while tackling something new, why not making a game instead?

Go is one of the languages that is known for making programming fun again, and what could be more fun than writing a game? Interested? Great, let's move on!

We will be writing a Pac Man clone for the terminal. While writing a game you are exposed to several interesting requirements that make a very good ground for exploring many features of the language, including input, output, functions, pointers, concurrency and a bit of math.

You also learn a bit more about the terminal and it's magical escape sequences.

## How to use this tutorial

This tutorial is design to teach the basic concepts of the Go programming language. You are expected to have a basic understanding on how programming languages work, as we are not going into details about the basics. Also, basic terminal knowledge is recommended as well.

## Compatibility

This tutorial has been tested on both Linux and Mac OSX environments. For Windows environments you may need to install a terminal emulator.

Please beware that since it relies on the terminal rendering the "graphics" it could produce different results for different kinds of terminals. If you have an issue feel free to raise it so we can find a proper solution, naming both your OS and terminal names and versions.

**Note:** it is a known issue that the terminal window on VS Code doesn't render the game correctly at this moment.

## Setup

In order to start, make sure you have Go installed in your system.

If you don't, please follow the instructions in [golang.org](https://golang.org)

## Step 00: Hello (Game) World

We are going to start by laying the ground a skeleton of what a game program looks like:

```go
package main

import "fmt"

func main() {
	// initialize game

	// load resources

	// game loop
	for {
		// process input

		// process movement

		// process collisions

		// update screen

		// check game over

		// Temp: break infinite loop
		fmt.Println("Hello, Pac Go!")
		break

		// repeat
	}
}
```

This code is also available as `main.go` in the root of this repository.

### Running your first Go program

Now that we have a `main.go` file (`.go` is the file extension for Go source code), you can run it by using the command line tool `go`. Just type:

```sh
$ go run main.go
Hello, Pac Go!
```

That's how we run a single file in Go. You can also build it as an executable program with the command `go build`. If you run `go build` it will compile the files in the current directory in a binary with the name of the directory. Then you can run it as a regular program, for example:

```sh
$ go build
$ ./pacgo
Hello, Pac Go!
```

For the purposes of this tutorial we are using only a single file of code (`main.go`), so you may use either `go build` and run the command or just `go run main.go` as it does both automatically.

### Understanding the program

Now let's have a close look of what the program does.

First line is the `package` name. Every valid Go program must have one of these. Also, if you want to make a **runnable** program you must have a package called `main` and a `main` function, which will be the entry point of the program.

We are creating an executable program so we are using `package main` on the top of the file.

Next are the `import` statements. You use those to make code in other packages accessible to your program.

Finally the `main` function. You define function in Go with the keyword `func` followed by it's name, it's parameters in between a pair of parenthesis, followed by the return value and finally the function body, which is contained in a pair of curly brackets `{}`. For example:

```go
func main() {

}
```

This is a function called main which takes no parameters and return nothing. That's the proper definition of a main function in Go, as opposed to the definitions other languages where the main function takes the command line arguments and return an integer value.

We have different ways to deal with the command line arguments and return values in Go, which we will see in Step 10.

In our main function we have some comments (any text after `//` is a comment) that are acting as placeholders for the actual game code. We will use those to drive each modification in a orderly way.

The most important concept in a game is what is called the game loop. That's basically an infinite loop where all the game mechanics are processed.

A loop in Go is defined with the keyword `for`. The `for` loop can have an initializer, a condition and a post processing step, but all of those are optional. If you omit all you have a loop that never ends:

```go
for {
    // I never end
}
```

We can exit an infinite loop with a `break` statement. We are using it in the sample code to end the infinite loop after printing "Hello, Pac Go!" with the `Println` function from the `fmt` package (comments omitted for brevity):

```go
func main() {
	for {
		fmt.Println("Hello, Pac Go!")
		break
    }
}
```

Of course, in this case the infinite loop with a non-conditional break is pointless, but it will make sense in the next steps!

[Take me to step 01!](step01/README.md)

## Contributing

This project is open source under the MIT license. If you want to contribute just submit a pull request.

If you are looking for inspiration you may browse the open issues or have a look at the [TODO](TODO.md) list. Everything on the TODO list should be planned as a new step on the tutorial unless otherwise noted.

## License

See [LICENSE](LICENSE)

## Contact Information

If you have any questions, please reach out to [daniela.petruzalek@gmail.com](mailto:daniela.petruzalek@gmail.com).