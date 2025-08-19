package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

// Screen holds the screen buffer and dimensions
type Screen struct {
	Width  int
	Height int
	Buffer bytes.Buffer
}

var screen Screen

func initialise() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	cbTerm.Run()
}

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	cookedTerm.Run()
}

// InitScreen initializes the screen buffer
func InitScreen(width, height int) {
	screen.Width = width
	screen.Height = height
	screen.Buffer.Grow((width + 1) * height)
}

// ClearScreen clears the screen buffer
func ClearScreen() {
	screen.Buffer.Reset()
	simpleansi.ClearScreen()
}

// Draw "draws" a string to the screen buffer
func Draw(row, col int, s string) {
	// For now, we'll just print directly to the screen.
	// We'll implement the buffer logic later.
	moveCursor(row, col)
	fmt.Print(s)
}

// Flush flushes the screen buffer to the terminal
func Flush() {
	fmt.Print(screen.Buffer.String())
}

func moveCursor(row, col int) {
	if cfg.UseEmoji {
		simpleansi.MoveCursor(row, col*2)
	} else {
		simpleansi.MoveCursor(row, col)
	}
}
