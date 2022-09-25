#!/bin/bash

INTALLATION_DIRECTORY="$HOME/Downloads/packages"
PACKAGE_URL="https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb"
PACKAGE_NAME="Google Chrome"
FILE_NAME="google-chrome.deb"

mkdir -p $INTALLATION_DIRECTORY
cd $INTALLATION_DIRECTORY

echo "Downloading $PACKAGE_NAME"

wget -c "$PACKAGE_URL" -P "$INTALLATION_DIRECTORY" -O "$FILE_NAME"

echo "Installing $PACKAGE_NAME"

sudo apt install "$FILE_NAME" -y
