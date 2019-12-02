package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

var (
	configFile = flag.String("config-file", "config.json", "path to custom configuration file")
	mazeFile   = flag.String("maze-file", "maze01.txt", "path to a custom maze file")
)

// Player is the player character \o/

type Point struct {
	row int
	col int
}

type Player struct {
	position Point
	origin   Point
}

var player Player

type GhostStatus string

const (
	Normal GhostStatus = "Normal"
	Blue   GhostStatus = "Blue"
)

// Ghost is the enemy that chases the player :O
type Ghost struct {
	position        Point
	initialPosition Point
	status          GhostStatus
}

var ghosts []*Ghost
var gostsStatusMx sync.RWMutex

// Config holds the emoji configuration
type Config struct {
	Player           string        `json:"player"`
	Ghost            string        `json:"ghost"`
	GhostBlue        string        `json:"ghost_blue"`
	Wall             string        `json:"wall"`
	Dot              string        `json:"dot"`
	Pill             string        `json:"pill"`
	Death            string        `json:"death"`
	Space            string        `json:"space"`
	UseEmoji         bool          `json:"use_emoji"`
	PillDurationSecs time.Duration `json:"pill_duration_secs"`
}

var cfg Config

func loadConfig() error {
	f, err := os.Open(*configFile)
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

func loadMaze() error {
	f, err := os.Open(*mazeFile)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = Player{position: Point{row, col}, origin: Point{row, col}}
			case 'G':
				ghosts = append(ghosts, &Ghost{Point{row, col}, Point{row, col}, "Normal"})
			case '.':
				numDots++
			}
		}
	}

	return nil
}

var maze []string
var score int
var numDots int
var lives = 3

func clearScreen() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	if cfg.UseEmoji {
		fmt.Printf("\x1b[%d;%df", row+1, col*2+1)
	} else {
		fmt.Printf("\x1b[%d;%df", row+1, col+1)
	}
}

func printScreen() {
	clearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Printf(cfg.Wall)
			case '.':
				fmt.Printf(cfg.Dot)
			case 'X':
				fmt.Printf(cfg.Pill)
			default:
				fmt.Printf(cfg.Space)
			}
		}
		fmt.Printf("\n")
	}

	moveCursor(player.position.row, player.position.col)
	fmt.Printf(cfg.Player)

	gostsStatusMx.RLock()
	for _, g := range ghosts {
		moveCursor(g.position.row, g.position.col)
		if g.status == Normal {
			fmt.Printf(cfg.Ghost)
		} else if g.status == Blue {
			fmt.Printf(cfg.GhostBlue)
		}
	}
	gostsStatusMx.RUnlock()

	moveCursor(len(maze)+1, 0)

	livesRemaining := strconv.Itoa(lives) //converts lives int to a string

	if cfg.UseEmoji {
		livesRemaining = getLivesAsEmoji()
	}

	fmt.Printf("Score: %v\tLives: %v\n", score, livesRemaining)
}

//concatenate the correct number of player emojis based on lives
func getLivesAsEmoji() string {
	buf := bytes.Buffer{}
	for i := lives; i > 0; i-- {
		buf.WriteString(cfg.Player)
	}
	return buf.String()
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze)-1 {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return
}

func movePlayer(dir string) {
	player.position.row, player.position.col = makeMove(player.position.row, player.position.col, dir)
	switch maze[player.position.row][player.position.col] {
	case '.':
		numDots--
		score++
		// Remove dot from the maze
		maze[player.position.row] = maze[player.position.row][0:player.position.col] + " " + maze[player.position.row][player.position.col+1:]
	case 'X':
		go proccessPill()
		// Remove pill from the maze
		maze[player.position.row] = maze[player.position.row][0:player.position.col] + " " + maze[player.position.row][player.position.col+1:]
	}
}

func updateGhosts(ghosts []*Ghost, ghostStatus GhostStatus) {
	gostsStatusMx.Lock()
	defer gostsStatusMx.Unlock()
	for _, g := range ghosts {
		g.status = ghostStatus
	}
}

var pillTimer *time.Timer

func proccessPill() {
	updateGhosts(ghosts, Blue)
	if pillTimer != nil {
		pillTimer.Stop()
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDurationSecs)
	pillTimer.Reset(time.Second * cfg.PillDurationSecs)
	<-pillTimer.C
	pillTimer.Stop()
	updateGhosts(ghosts, Normal)
}

func drawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}

func moveGhosts() {
	for _, g := range ghosts {
		dir := drawDirection()
		g.position.row, g.position.col = makeMove(g.position.row, g.position.col, dir)
	}
}

func initialize() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
	}
}

func main() {
	flag.Parse()

	// initialize game
	initialize()
	defer cleanup()

	// load resources
	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

	err = loadConfig()
	if err != nil {
		log.Printf("Error loading configuration: %v\n", err)
		return
	}

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
		// process movement
		select {
		case inp := <-input:
			if inp == "ESC" {
				lives = 0
			}
			movePlayer(inp)
		default:
		}

		moveGhosts()

		// process collisions
		for _, g := range ghosts {
			if player.position.row == g.position.row && player.position.col == g.position.col {
				gostsStatusMx.RLock()
				if g.status == Normal {
					lives = lives - 1
					if lives != 0 {
						moveCursor(player.position.row, player.position.col)
						fmt.Printf(cfg.Death)
						moveCursor(len(maze)+2, 0)
						gostsStatusMx.RUnlock()
						updateGhosts(ghosts, Normal)
						time.Sleep(1000 * time.Millisecond) //dramatic pause before reseting player position
						player.position = player.origin
					}
				} else if g.status == Blue {
					gostsStatusMx.RUnlock()
					updateGhosts([]*Ghost{g}, Normal)
					g.position.row, g.position.col = g.initialPosition.row, g.initialPosition.col
				}
			}
		}

		// update screen
		printScreen()

		// check game over
		if numDots == 0 || lives == 0 {
			if lives == 0 {
				moveCursor(player.position.row, player.position.col)
				fmt.Printf(cfg.Death)
				moveCursor(player.origin.row, player.origin.col-1)
				fmt.Printf("GAME OVER")
				moveCursor(len(maze)+2, 0)
			}
			break
		}

		// repeat
		time.Sleep(200 * time.Millisecond)
	}
}
