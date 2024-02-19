# GitMan

GitMan is a simple package manager written in Go. It allows you to install packages from a centralized repository using Git.

## Usage

To use GitMan, follow these steps:

1. Download the GitMan binary or build it from the source:

    ```bash
    make build
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