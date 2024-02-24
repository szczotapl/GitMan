![GitHub License](https://img.shields.io/github/license/riviox/Gitman) ![GitHub repo size](https://img.shields.io/github/repo-size/riviox/GitMan) ![GitHub Repo stars](https://img.shields.io/github/stars/riviox/GitMan) 


______________________
# GitMan

GitMan is a simple package manager written in Go. It allows you to install packages from a centralized repository using Git.

## Usage

To use GitMan, follow these steps:

1. Install GitMan with:
    ```
    curl -sSL https://raw.githubusercontent.com/riviox/GitMan/main/install.sh | bash
    ```
    * Dependencies:
        - Go
        - Git
          - Install deps with pacman:
            ```
            sudo pacman -S go git
            ```
    * To update GitMan run:
    ```
    gitman_update
    ```

2. Install a package:

    ```bash
    gitman -S <package_name>
    ```

    Replace `<package_name>` with the name of the package you want to install.

3. List packages:
    ```
    gitman -L
    ```
3. Remove package:
    ```
    gitman -R <package_name>
    ```
    Replace `<package_name>` with the name of the package you want to remove.

## Example

```bash
# To install 'hellocpp'
gitman -S hellocpp
# To remove 'hellocpp'
gitman -R hellocpp
```
