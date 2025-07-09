//go:build windows

package main

import (
	"bufio"
	"fmt"
	"os"
)

// waitForExit prompts the user to press Enter before exiting.
// This is only included in Windows builds to prevent the console window
// from closing immediately when the executable is double-clicked.
func waitForExit() {
	fmt.Println("\nPress Enter to exit.")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}