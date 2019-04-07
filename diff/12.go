33a34,48
> func readInput() (string, error) {
> 	buffer := make([]byte, 100)
> 
> 	cnt, err := os.Stdin.Read(buffer)
> 	if err != nil {
> 		return "", err
> 	}
> 
> 	if cnt == 1 && buffer[0] == 0x1b {
> 		return "ESC", nil
> 	}
> 
> 	return "", nil
> }
> 
46a62,66
> 		input, err := readInput()
> 		if err != nil {
> 			log.Printf("Error reading input: %v", err)
> 			break
> 		}
56,58c76,78
< 
< 		// Temp: break infinite loop
< 		break
---
> 		if input == "ESC" {
> 			break
> 		}
