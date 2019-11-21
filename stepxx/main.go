package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/danicat/simpleansi"
)

var (
	configFile = flag.String("config-file", "config.json", "path to custom configuration file")
	mazeFile   = flag.String("maze-file", "maze01.txt", "path to a custom maze file")
)

var player *Player
var sprites []Sprite

// Config holds the emoji configuration
type Config struct {
	Player    string        `json:"player"`
	Ghost     string        `json:"ghost"`
	Wall      string        `json:"wall"`
	Dot       string        `json:"dot"`
	Pill      string        `json:"pill"`
	Death     string        `json:"death"`
	Space     string        `json:"space"`
	Chaser    string        `json:"chaser"`
	UseEmoji  bool          `json:"use_emoji"`
	FrameRate time.Duration `json:"frame_rate"`
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

	// Default: 5 FPS
	if cfg.FrameRate == 0 {
		cfg.FrameRate = 5
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
				player = NewPlayer(row, col, 1, cfg.Player)
				sprites = append(sprites, player)
			case 'G':
				sprites = append(sprites, NewGhost(row, col, cfg.Ghost))
			case 'C':
				sprites = append(sprites, NewChaser(row, col, cfg.Chaser))
			case '.':
				numDots++
			}
		}
	}

	return nil
}

var maze []string
var numDots int

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
			case '.':
				fmt.Print(cfg.Dot)
			case 'X':
				fmt.Print(cfg.Pill)
			default:
				fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}

	for _, s := range sprites {
		moveCursor(s.Pos())
		fmt.Print(s.Img())
	}

	moveCursor(len(maze)+1, 0)
	fmt.Println("Score:", player.score, "\tLives:", player.lives)
}

func main() {
	flag.Parse()

	// initialise game
	initialise()
	defer cleanup()

	// load resources
	err := loadConfig()
	if err != nil {
		log.Println("Error loading configuration:", err)
		return
	}

	err = loadMaze()
	if err != nil {
		log.Println("Error loading maze:", err)
		return
	}

	// game loop
	for {
		// process movement
		for _, s := range sprites {
			go s.Move()
		}

		// update screen
		printScreen()

		// check game over
		if numDots == 0 || player.lives <= 0 {
			if player.lives <= 0 {
				moveCursor(player.Pos())
				fmt.Print(cfg.Death)
				moveCursor(len(maze)+2, 0)
			}
			break
		}

		// wait before rendering next frame
		time.Sleep(1000 / cfg.FrameRate * time.Millisecond)
	}
}
