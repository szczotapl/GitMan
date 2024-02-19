package main

import (
	"encoding/json"
	"flag"
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

const (
	defaultInstallDir = ".gitman/packages"
	defaultJSONURL    = "https://raw.githubusercontent.com/riviox/GitMan/main/packages.json"
)

var (
	installDir string
	jsonURL    string
	listFlag   bool
)

func init() {
	flag.StringVar(&installDir, "install-dir", defaultInstallDir, "Specify the installation directory")
	flag.StringVar(&jsonURL, "json-url", defaultJSONURL, "Specify the URL for the packages.json file")
	flag.BoolVar(&listFlag, "L", false, "List available packages")
	flag.Parse()
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

func listPackages() {
	jsonContent, err := downloadFile(jsonURL)
	if err != nil {
		fmt.Println("Error downloading packages.json:", err)
		os.Exit(1)
	}

	packages, err := parsePackagesJSON(jsonContent)
	if err != nil {
		fmt.Println("Error parsing packages.json:", err)
		os.Exit(1)
	}

	fmt.Println("Available Packages:")
	for _, pkg := range packages {
		fmt.Printf("- %s\n", pkg.Name)
	}
}

func main() {
	args := os.Args

	if listFlag {
		listPackages()
		return
	}

	if len(args) < 2 {
		fmt.Println("Usage: gitman.go <options>")
		flag.PrintDefaults()
		os.Exit(1)
	}


	switch args[1] {
	case "-S":
		if len(args) != 3 {
			fmt.Println("Usage: gitman.go -S <package_name>")
			os.Exit(1)
		}

		packageName := args[2]

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting user home directory:", err)
			os.Exit(1)
		}

		installDir := filepath.Join(homeDir, installDir)

		jsonContent, err := downloadFile(jsonURL)
		if err != nil {
			fmt.Println("Error downloading packages.json:", err)
			os.Exit(1)
		}

		packages, err := parsePackagesJSON(jsonContent)
		if err != nil {
			fmt.Println("Error parsing packages.json:", err)
			os.Exit(1)
		}

		var selectedPackage Package
		for _, pkg := range packages {
			if pkg.Name == packageName {
				selectedPackage = pkg
				break
			}
		}

		if selectedPackage.Name != "" {
			fmt.Println("Package found:", selectedPackage.Name)

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
	case "-L":
		listPackages()
	default:
		fmt.Println("Unknown option:", args[1])
		os.Exit(1)
	}
}
