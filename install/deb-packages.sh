#!/bin/bash

INTALLATION_DIRECTORY="$HOME/Downloads/packages"

declare -a DEB_PACKAGES=("https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb" "https://updates.insomnia.rest/downloads/ubuntu/latest")

echo "Download .deb packages"

for i in "${DEB_PACKAGES[@]}"; do
    wget -c "$i" -P "$INTALLATION_DIRECTORY"
done

echo "Installing .deb packages"

for i in "$INTALLATION_DIRECTORY"/*; do
    sudo apt install "$i" -y
done
