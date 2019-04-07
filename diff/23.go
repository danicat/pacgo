7a8
> 	"os/exec"
27a29,37
> func clearScreen() {
> 	fmt.Printf("\x1b[2J")
> 	moveCursor(0, 0)
> }
> 
> func moveCursor(row, col int) {
> 	fmt.Printf("\x1b[%d;%df", row+1, col+1)
> }
> 
28a39
> 	clearScreen()
48a60,79
> func initialize() {
> 	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
> 	cbTerm.Stdin = os.Stdin
> 
> 	err := cbTerm.Run()
> 	if err != nil {
> 		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
> 	}
> }
> 
> func cleanup() {
> 	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
> 	cookedTerm.Stdin = os.Stdin
> 
> 	err := cookedTerm.Run()
> 	if err != nil {
> 		log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
> 	}
> }
> 
50a82,83
> 	initialize()
> 	defer cleanup()
