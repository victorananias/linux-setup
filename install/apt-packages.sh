#!/bin/bash

declare -a APT_PACKAGES=(
    "firmware-linux-nonfree"
    "curl"
    "zsh"
    "git"
    "fonts-firacode"
    "default-jre"
    "default-jdk"
    "calibre"
    "qbittorrent"
    "flatpak"
    "pulseeffects"
    "gnome-tweak-tool"
    "keepassxc"
    "xclip"
)

echo "Install apt packages"

for i in "${APT_PACKAGES[@]}"; do
    sudo apt install -y $i
done
