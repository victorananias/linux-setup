#!/bin/bash

DOWNLOADS="Installation Files"
URL_GOOGLE_CHROME="https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb"
URL_UNITY_HUB="https://public-cdn.cloud.unity3d.com/hub/prod/UnityHub.AppImage"

# PACKAGES
declare -a APT_PACKAGES=("curl" "snapd" "zsh" "git" "fonts-firacode" "vlc" "steam" "remmina" "remmina-plugin-rdp" "php" "php-cli" "php-fpm" "php-json" "php-pdo" "php-mysql" "php-zip" "php-gd" "php-mbstring" "php-curl" "php-xml" "php-pear" "php-bcmath")
declare -a SNAP_PACKAGES=("spotify code  krita blender")

# UPDATE
sudo apt update -y
sudo apt upgrade -y

for i in "${APT_PACKAGES[@]}"; do
    sudo apt install -y $i
done

# install snaps
for i in "${SNAP_PACKAGES[@]}"; do
    sudo snap install $i
done

cd /home/$USER/Downloads
mkdir "$DOWNLOADS"
wget -c "$URL_UNITY_HUB" -P "$DOWNLOADS"

wget -c "$URL_GOOGLE_CHROME" -P "$DOWNLOADS"
sudo apt install "$DOWNLOADS/google-chrome-stable_current_amd64.deb"
