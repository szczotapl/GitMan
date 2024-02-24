package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	updateCommands := []string{
		"rm -rf ~/.gitman/src",
		"curl -sSL https://raw.githubusercontent.com/riviox/GitMan/main/install.sh | bash",
	}

	for _, cmdStr := range updateCommands {
		cmd := exec.Command("bash", "-c", cmdStr)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running update command '%s': %s\n", cmdStr, err)
			os.Exit(1)
		}
	}

	fmt.Println("Update completed successfully!")
}
