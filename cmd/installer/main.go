package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the NEOS Manual Install Helper for the comma two.")
	fmt.Println("This tool provides a workaround to bypass the NEOS setup screen.")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("First, ensure your comma two is on the same Wi-Fi network as this computer.")
	fmt.Println("On your device, go to Settings -> Wi-Fi -> More Options -> Advanced.")
	fmt.Println("You will find the IPv4 address there.")
	fmt.Println("It will likely start with 192.168.x.x, 10.x.x.x, or 172.16.x.x.")
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

	waitForExit()
}
