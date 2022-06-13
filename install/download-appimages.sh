#!/bin/bash

APPIMAGES_DIRECTORY="$HOME/Applications"

declare -a APPIMAGES=()

echo "Download AppImages"

for i in "${APPIMAGES[@]}"
do
    wget -c "$i" -P "$APPIMAGES_DIRECTORY"
done
