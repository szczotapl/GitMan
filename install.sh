#!/bin/bash

print_error() {
    echo -e "\033[0;31mError: $1\033[0m" >&2
    exit 1
}

print_info() {
    echo -e "\033[0;32m$1\033[0m"
}

gitman_dir="$HOME/.gitman/src"

if [ -d "$gitman_dir" ]; then
    print_info "GitMan directory already exists at $gitman_dir."
else
    git clone https://github.com/riviox/GitMan.git "$gitman_dir" --depth 1 || print_error "Failed to clone the GitMan repository."
    print_info "GitMan repository cloned to $gitman_dir."
fi

cd "$gitman_dir"
sudo make install || print_error "Failed to install GitMan."
print_info "GitMan installed successfully!"
