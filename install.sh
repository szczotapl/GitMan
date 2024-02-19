#!/bin/bash

git clone https://github.com/riviox/GitMan.git ~/.gitman/src

if [ $? -ne 0 ]; then
    echo "Error: Failed to clone the GitMan repository."
    exit 1
fi
echo "The GitMan repository has been successfully cloned to ~/.gitman/src."
cd ~/.gitman/src
make install
echo "Successfully installed GitMan!"