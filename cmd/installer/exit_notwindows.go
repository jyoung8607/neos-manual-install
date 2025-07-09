//go:build !windows

package main

// waitForExit is a no-op on non-Windows platforms.
// The "press enter to exit" behavior is only needed on Windows
// when the executable is run by double-clicking.
func waitForExit() {
	// Do nothing
}