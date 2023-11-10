// Package helpers provides utilities for interacting with the terminal and formatting output.
package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// Terminal struct holds information about the user's terminal environment.
type Terminal struct {
	Height                  int
	Width                   int
	OutputType              string
	NumberOfSupportedColors int
	TERM                    string
	SHELL                   string
	COLORTERM               string
}

// ErrorLevel type for defining constants for the error levels
type ErrorLevel int

// Define constants for the error levels
const (
	Alert ErrorLevel = iota
	Critical
	Error
	Warn
	Notice
	Info
	Debug
	LightPurple
	Teal
	DarkGreen
	Brown
	Cyan
)

// colorMap maps error levels to their respective ANSI color codes
var colorMap = map[ErrorLevel]string{
	Alert:       "\033[38;5;201m", // Magenta
	Critical:    "\033[38;5;214m", // Orange
	Error:       "\033[38;5;196m", // Light Red
	Warn:        "\033[38;5;226m", // Yellow
	Notice:      "\033[38;5;117m", // Light Blue
	Info:        "\033[38;5;250m", // Gray
	Debug:       "\033[38;5;120m", // Light Green
	LightPurple: "\033[38;5;141m", // Light Purple
	Teal:        "\033[38;5;49m",  // Teal
	DarkGreen:   "\033[38;5;34m",  // Dark Green
	Brown:       "\033[38;5;130m", // Brown
	Cyan:        "\033[38;5;51m",  // Cyan
}

// ClearTerminal clears the terminal screen based on the operating system.
func ClearTerminal() error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// TerminalInfo prints collects various information about the current terminal.
func TerminalInfo() (*Terminal, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, fmt.Errorf("TerminalInfo(): Error getting terminal size: %w", err)
	}

	outputType := "Not a terminal"
	if term.IsTerminal(int(os.Stdout.Fd())) {
		outputType = "Terminal"
	}

	colorsOutput, err := exec.Command("tput", "colors").Output()
	if err != nil {
		return nil, fmt.Errorf("TerminalInfo(): Error getting number of supported colors: %w", err)
	}
	numberOfSupportedColors, err := strconv.Atoi(strings.TrimSpace(string(colorsOutput)))
	if err != nil {
		return nil, fmt.Errorf("TerminalInfo(): Error converting number of colors to int: %w", err)
	}

	terminalInfo := Terminal{
		Height:                  height,
		Width:                   width,
		OutputType:              outputType,
		NumberOfSupportedColors: numberOfSupportedColors,
		TERM:                    os.Getenv("TERM"),
		SHELL:                   os.Getenv("SHELL"),
		COLORTERM:               os.Getenv("COLORTERM"),
	}

	return &terminalInfo, nil
}

// TerminalColor prints the given string to the terminal in the color corresponding to the error level
func TerminalColor(message string, level ErrorLevel) {
	colorCode, ok := colorMap[level]
	if !ok {
		fmt.Println(message)
		return
	}
	fmt.Printf("%s%s\033[0m\n", colorCode, message)
}
