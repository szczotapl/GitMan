package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Package struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

func downloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parsePackagesJSON(jsonContent []byte) ([]Package, error) {
	var packages []Package
	err := json.Unmarshal(jsonContent, &packages)
	if err != nil {
		return nil, err
	}
	return packages, nil
}

func cloneGitRepository(repository, installDir string) error {
	err := os.MkdirAll(installDir, os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "clone", repository, installDir)
	err = cmd.Run()
	if err != nil {
		return err
	}

	installCmd := exec.Command("make", "install")
	installCmd.Dir = installDir
	return installCmd.Run()
}

func main() {
	args := os.Args

	if len(args) != 3 || args[1] != "-S" {
		fmt.Println("Usage: gitman.go -S package_name")
		os.Exit(1)
	}

	packageName := args[2]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}

	installDir := filepath.Join(homeDir, ".gitman/packages")

	// Pobierz zawartość packages.json z podanego URL
	jsonContent, err := downloadFile("https://raw.githubusercontent.com/riviox/GitMan/main/packages.json")
	if err != nil {
		fmt.Println("Error downloading packages.json:", err)
		os.Exit(1)
	}

	// Analizuj packages.json i uzyskaj listę pakietów
	packages, err := parsePackagesJSON(jsonContent)
	if err != nil {
		fmt.Println("Error parsing packages.json:", err)
		os.Exit(1)
	}

	// Sprawdź, czy pakiet o podanej nazwie istnieje
	var selectedPackage Package
	for _, pkg := range packages {
		if pkg.Name == packageName {
			selectedPackage = pkg
			break
		}
	}

	if selectedPackage.Name != "" {
		fmt.Println("Package found:", selectedPackage.Name)

		// Zrealizuj kroki instalacyjne
		err := cloneGitRepository(selectedPackage.Repository, installDir)
		if err != nil {
			fmt.Println("Error cloning repository or making install:", err)
			os.Exit(1)
		}

		fmt.Println("Package installed successfully!")
	} else {
		fmt.Println("Package not found:", packageName)
		os.Exit(1)
	}
}
