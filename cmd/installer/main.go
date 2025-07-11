package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fork struct {
	Name   string
	Owner  string
	Branch string
}

var forks = []Fork{
	{Name: "Stock", Owner: "commaai", Branch: "release2"},
	{Name: "Dragonpilot", Owner: "dragonpilot-community", Branch: "r2"},
}

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

	fmt.Println("\nAvailable forks:")
	for i, fork := range forks {
		fmt.Printf("%d. %s (%s/%s)\n", i+1, fork.Name, fork.Owner, fork.Branch)
	}
	fmt.Printf("%d. Custom\n", len(forks)+1)

	fmt.Print("Select a fork to install: ")
	response, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(response))

	var githubOwner, githubBranch string
	if choice > 0 && choice <= len(forks) {
		selectedFork := forks[choice-1]
		githubOwner = selectedFork.Owner
		githubBranch = selectedFork.Branch
	} else {
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
