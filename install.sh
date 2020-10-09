#!/bin/bash

DOWNLOADS="$HOME/Downloads/Installation Files"
APPIMAGES_DIRECTORY="$HOME/Applications"
PROGRAMS="./programs"

# PACKAGES
declare -a DEB_PACKAGES=("https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb")
declare -a APT_PACKAGES=("curl" "snapd" "zsh" "git" "fonts-firacode" "vlc" "steam" "remmina" "remmina-plugin-rdp" "php" "php-cli" "php-fpm" "php-json" "php-pdo" "php-mysql" "php-zip" "php-gd" "php-mbstring" "php-curl" "php-xml" "php-pear" "php-bcmath")
declare -a SNAP_PACKAGES=("insomnia" "krita" "blender --classic")
declare -a APPIMAGES=("https://public-cdn.cloud.unity3d.com/hub/prod/UnityHub.AppImage")

# UPDATE
sudo apt update -y
sudo apt upgrade -y

# install apt packages
for i in "${APT_PACKAGES[@]}"
do
    sudo apt install -y $i
done

# install snap packages
for i in "${SNAP_PACKAGES[@]}"
do
    sudo snap install $i
done

# download deb packages
for i in "${DEB_PACKAGES[@]}"
do
    wget -c "$i" -P "$DOWNLOADS"
done

# install deb packages
for i in "$DOWNLOADS"/*
do
    sudo apt install "$i" -y
done

# download AppImages
for i in "${APPIMAGES[@]}"
do
    wget -c "$i" -P "$APPIMAGES_DIRECTORY"
done

# install deb packages
for i in "$PROGRAMS"/*
do
    source $i
done