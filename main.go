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
	setConsoleColors() // init of terminal
	code := strconv.Itoa(colorcode)
	color := "\u001b[38;5;" + code + "m" // acii code used by terminals to print colors
	return fmt.Sprintf(color)

}

func main() {
	fmt.Printf("%vThis is color\n", Color(140)) // color code 140
	fmt.Print("\u001b[0m") // This is to reset the terminal color 
}
