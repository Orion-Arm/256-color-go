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
	fmt.Print("\u001b[0m")
	return fmt.Sprintf(color)

}

func main() {
	fmt.Printf("%vThis is color\n", Color(140))
	fmt.Print("\u001b[0m")
}
