#!/bin/bash

declare -a APT_PACKAGES=("curl" "snapd"  "zsh" "git" "fonts-firacode" "default-jre" "default-jdk" "calibre" "qbittorrent")

echo "Install apt packages"

for i in "${APT_PACKAGES[@]}"
do
    sudo apt install -y $i
done
