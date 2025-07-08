package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the NEOS openpilot installer.")
	fmt.Println("This tool will guide you through installing openpilot on your comma device.")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("First, please ensure your device is connected to Wi-Fi.")
	fmt.Println("On your device, go to Settings -> Wi-Fi -> More Options -> Advanced.")
	fmt.Println("You will find the IPv4 address there (e.g., 192.168.1.100).")
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Print("Enter the device IP address: ")
	ipAddress, _ := reader.ReadString('\n')
	ipAddress = strings.TrimSpace(ipAddress)

	fmt.Print("Install a custom fork? (y/N): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	githubOwner := "commaai"
	githubBranch := "release2"

	if response == "y" {
		fmt.Print("Enter the GitHub repository owner: ")
		owner, _ := reader.ReadString('\n')
		githubOwner = strings.TrimSpace(owner)

		fmt.Print("Enter the GitHub branch name: ")
		branch, _ := reader.ReadString('\n')
		githubBranch = strings.TrimSpace(branch)
	}

	fmt.Println("\n-----------------------------------------------------------------------")
	fmt.Printf("Device IP: %s\n", ipAddress)
	fmt.Printf("GitHub Repo: https://github.com/%s/openpilot.git\n", githubOwner)
	fmt.Printf("Branch: %s\n", githubBranch)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("\nConnecting to device...")
	client, err := connect(ipAddress)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	if err := runSetup(client, githubOwner, githubBranch); err != nil {
		fmt.Printf("An error occurred during setup: %v\n", err)
	} else {
		fmt.Println("Installation script finished successfully.")
	}

	fmt.Println("\nPress Enter to exit.")
	reader.ReadString('\n')
}