#!/bin/bash

INTALLATION_DIRECTORY="$HOME/Downloads/packages"
PACKAGE_URL="https://updates.insomnia.rest/downloads/ubuntu/latest"
PACKAGE_NAME="Insomnia"
FILE_NAME="insomnia.deb"

mkdir -p $INTALLATION_DIRECTORY
cd $INTALLATION_DIRECTORY

echo "Downloading $PACKAGE_NAME"

wget -c "$PACKAGE_URL" -P "$INTALLATION_DIRECTORY" -O "$FILE_NAME"

echo "Installing $PACKAGE_NAME"

sudo apt install "$FILE_NAME" -y
