# GitMan

GitMan is a simple package manager written in Go. It allows you to install packages from a centralized repository using Git.

## Usage

To use GitMan, follow these steps:

1. Install GitMan with:
    ```
    curl -sSL https://raw.githubusercontent.com/riviox/GitMan/main/install.sh | bash
    ```

2. Install a package:

    ```bash
    ./gitman -S <package_name>
    ```

    Replace `<package_name>` with the name of the package you want to install.

## Example

```bash
./gitman -S gfetch
```