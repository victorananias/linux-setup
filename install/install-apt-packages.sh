#!/bin/bash

declare -a APT_PACKAGES=("curl" "zsh" "git" "fonts-firacode" "default-jre" "default-jdk" "calibre" "qbittorrent" "flatpak")

echo "Install apt packages"

for i in "${APT_PACKAGES[@]}"; do
    sudo apt install -y $i
done
