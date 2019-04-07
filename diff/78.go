4a5
> 	"encoding/json"
28a30,59
> // Config holds the emoji configuration
> type Config struct {
> 	Player   string `json:"player"`
> 	Ghost    string `json:"ghost"`
> 	Wall     string `json:"wall"`
> 	Dot      string `json:"dot"`
> 	Pill     string `json:"pill"`
> 	Death    string `json:"death"`
> 	Space    string `json:"space"`
> 	UseEmoji bool   `json:"use_emoji"`
> }
> 
> var cfg Config
> 
> func loadConfig() error {
> 	f, err := os.Open("config.json")
> 	if err != nil {
> 		return err
> 	}
> 	defer f.Close()
> 
> 	decoder := json.NewDecoder(f)
> 	err = decoder.Decode(&cfg)
> 	if err != nil {
> 		return err
> 	}
> 
> 	return nil
> }
> 
69c100,104
< 	fmt.Printf("\x1b[%d;%df", row+1, col+1)
---
> 	if cfg.UseEmoji {
> 		fmt.Printf("\x1b[%d;%df", row+1, col*2+1)
> 	} else {
> 		fmt.Printf("\x1b[%d;%df", row+1, col+1)
> 	}
78c113
< 				fallthrough
---
> 				fmt.Printf(cfg.Wall)
80c115
< 				fmt.Printf("%c", chr)
---
> 				fmt.Printf(cfg.Dot)
82c117
< 				fmt.Printf(" ")
---
> 				fmt.Printf(cfg.Space)
89c124
< 	fmt.Printf("P")
---
> 	fmt.Printf(cfg.Player)
93c128
< 		fmt.Printf("G")
---
> 		fmt.Printf(cfg.Ghost)
222a258,263
> 	err = loadConfig()
> 	if err != nil {
> 		log.Printf("Error loading configuration: %v\n", err)
> 		return
> 	}
> 
261a303,307
> 			if lives == 0 {
> 				moveCursor(player.row, player.col)
> 				fmt.Printf(cfg.Death)
> 				moveCursor(len(maze)+2, 0)
> 			}
