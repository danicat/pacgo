9a10
> 	"time"
221a223,235
> 	// process input (async)
> 	input := make(chan string)
> 	go func(ch chan<- string) {
> 		for {
> 			input, err := readInput()
> 			if err != nil {
> 				log.Printf("Error reading input: %v", err)
> 				ch <- "ESC"
> 			}
> 			ch <- input
> 		}
> 	}(input)
> 
224,228c238,245
< 		// process input
< 		input, err := readInput()
< 		if err != nil {
< 			log.Printf("Error reading input: %v", err)
< 			break
---
> 		// process movement
> 		select {
> 		case inp := <-input:
> 			if inp == "ESC" {
> 				lives = 0
> 			}
> 			movePlayer(inp)
> 		default:
231,232d247
< 		// process movement
< 		movePlayer(input)
246c261
< 		if input == "ESC" || numDots == 0 || lives == 0 {
---
> 		if numDots == 0 || lives == 0 {
250a266
> 		time.Sleep(200 * time.Millisecond)
