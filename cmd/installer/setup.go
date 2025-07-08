package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func runSetup(client *ssh.Client, githubOwner, githubBranch string) error {
	fmt.Println("Starting setup on the device...")

	// 1. Change to /data directory
	fmt.Println("Changing to /data directory...")
	if _, err := executeCommand(client, "cd /data"); err != nil {
		// This command will likely return an empty output on success, so we ignore it.
		// A failure would be caught by the error.
	}

	// 2. Remove existing openpilot directory
	fmt.Println("Removing existing openpilot directory...")
	if _, err := executeCommand(client, "rm -rf openpilot"); err != nil {
		return fmt.Errorf("failed to remove openpilot directory: %v", err)
	}

	// 3. Clone openpilot repository
	fmt.Println("Cloning openpilot repository...")
	cloneCmd := fmt.Sprintf("git clone https://github.com/%s/openpilot.git openpilot -b %s --recurse-submodules --depth 1", githubOwner, githubBranch)
	if _, err := executeCommand(client, "cd /data && "+cloneCmd); err != nil {
		return fmt.Errorf("failed to clone openpilot: %v", err)
	}

	// 4. Create continue.sh script
	fmt.Println("Creating continue.sh script...")
	continueScript := `#!/usr/bin/bash\n\ncd /data/openpilot\n./launch_openpilot.sh\n`
	createScriptCmd := fmt.Sprintf(`echo $'%s' > /data/data/com.termux/files/continue.sh`, continueScript)
	if _, err := executeCommand(client, createScriptCmd); err != nil {
		return fmt.Errorf("failed to create continue.sh: %v", err)
	}

	// 5. Make continue.sh executable
	fmt.Println("Making continue.sh executable...")
	if _, err := executeCommand(client, "chmod +x /data/data/com.termux/files/continue.sh"); err != nil {
		return fmt.Errorf("failed to make continue.sh executable: %v", err)
	}

	// 6. Reboot the device
	fmt.Println("Setup complete. Rebooting device...")
	if _, err := executeCommand(client, "reboot"); err != nil {
		// The reboot command might close the connection before a response is received.
		// We can consider this a success if there's no immediate error.
		fmt.Println("Reboot command sent. The device is now restarting.")
	}

	return nil
}