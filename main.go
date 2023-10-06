// modified by Abhishek Mishra 
// can print 612 colors now
package main

import (
	"fmt"
	"strconv"

	"golang.org/x/sys/windows"
)

func setConsoleColors() error {
	console := windows.Stdout
	var consoleMode uint32
	windows.GetConsoleMode(console, &consoleMode)
	consoleMode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	return windows.SetConsoleMode(console, consoleMode)
}

func Color(colorcode int) string {
	setConsoleColors()
	code := strconv.Itoa(colorcode)
	color := "\u001b[38;5;" + code + "m" 
	return fmt.Sprintf(color)
}

func main() {
	for colorCode := 1; colorCode <= 612; colorCode++ {
		fmt.Printf("%vColor %d\n", Color(colorCode), colorCode)
		fmt.Print("\u001b[0m") 
	}
}
