#!/bin/bash

INTALLATION_DIRECTORY="$HOME/Downloads/packages"
PACKAGE_URL="https://storage.googleapis.com/golang/getgo/installer_linux"
PACKAGE_NAME="Go Installer"
FILE_NAME="go-installer"

mkdir -p $INTALLATION_DIRECTORY
cd $INTALLATION_DIRECTORY

echo "Downloading $PACKAGE_NAME"

wget -c "$PACKAGE_URL" -P "$INTALLATION_DIRECTORY" -o "$FILE_NAME"

sudo chmod +x "$FILE_NAME"

echo "Installing $PACKAGE_NAME"

bash $FILE_NAME

mkdir ~/go

rm $FILE_NAME
