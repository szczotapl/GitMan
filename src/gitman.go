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

	"github.com/fatih/color"
)

type Package struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

const (
	defaultInstallDir  = ".gitman/packages"
	defaultSourcesFile = ".gitman/config/sources.json"
	defaultJSONURL     = "https://raw.githubusercontent.com/riviox/GitMan/main/packages.json"
)

var (
	installDir  string
	sourcesFile string
	listFlag    bool
	packageName string
)

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

func cloneGitRepository(repository, installDir, packageName string) error {
	packageDir := filepath.Join(installDir, packageName)

	err := os.MkdirAll(installDir, os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "clone", repository, packageDir)
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = os.Chdir(packageDir)
	if err != nil {
		return err
	}

	installCmd := exec.Command("make", "install")
	err = installCmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func uninstallPackage(packageName, installDir string) error {
	packageDir := filepath.Join(installDir, packageName)

	if _, err := os.Stat(packageDir); os.IsNotExist(err) {
		return fmt.Errorf("Package not found: %s", packageName)
	}

	err := os.Chdir(packageDir)
	if err != nil {
		return err
	}

	uninstallCmd := exec.Command("make", "uninstall")

	output, err := uninstallCmd.CombinedOutput()
	if err != nil {
		color.Red("Error uninstalling package: %s\n", err)
		fmt.Printf("Command output:\n%s\n", output)
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	err = os.RemoveAll(packageDir)
	if err != nil {
		return err
	}

	return nil
}

func listPackages() {
	jsonContent, err := downloadFile(defaultJSONURL)
	if err != nil {
		color.Red("Error downloading packages.json: %s\n", err)
		os.Exit(1)
	}

	packages, err := parsePackagesJSON(jsonContent)
	if err != nil {
		color.Red("Error parsing packages.json: %s\n", err)
		os.Exit(1)
	}

	color.Green("Available Packages:\n")
	for _, pkg := range packages {
		color.Cyan("- %s\n", pkg.Name)
	}
}

func main() {
	flag.StringVar(&installDir, "install-dir", defaultInstallDir, "Specify the installation directory")
	flag.StringVar(&sourcesFile, "sources-file", defaultSourcesFile, "Specify the sources configuration file")
	flag.BoolVar(&listFlag, "L", false, "List available packages")
	flag.StringVar(&packageName, "S", "", "Specify the package to install")

	uninstallFlag := flag.String("R", "", "Specify the package to uninstall")

	flag.Parse()

	if listFlag {
		listPackages()
		return
	}

	if packageName != "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			color.Red("Error getting user home directory: %s\n", err)
			os.Exit(1)
		}

		installDir := filepath.Join(homeDir, installDir)

		jsonContent, err := downloadFile(defaultJSONURL)
		if err != nil {
			color.Red("Error downloading packages.json: %s\n", err)
			os.Exit(1)
		}

		packages, err := parsePackagesJSON(jsonContent)
		if err != nil {
			color.Red("Error parsing packages.json: %s\n", err)
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
			color.Green("Package found: %s\n", selectedPackage.Name)

			err := cloneGitRepository(selectedPackage.Repository, installDir, packageName)
			if err != nil {
				color.Red("Error cloning repository or making install: %s\n", err)
				os.Exit(1)
			}
			color.Green("Package '%s' installed successfully!\n", packageName)
		} else {
			color.Red("Package not found: %s\n", packageName)
			os.Exit(1)
		}
	} else if *uninstallFlag != "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			color.Red("Error getting user home directory: %s\n", err)
			os.Exit(1)
		}

		installDir := filepath.Join(homeDir, installDir)

		err = uninstallPackage(*uninstallFlag, installDir)
		if err != nil {
			color.Red("Error uninstalling package: %s\n", err)
			os.Exit(1)
		}

		color.Green("Package uninstalled successfully!\n")
	} else {
		color.Red("Usage: gitman <options>\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
