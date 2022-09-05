#!/bin/bash

declare -a FLATPAKS=("com.discordapp.Discord" "org.mozilla.firefox")

echo "Install flatpaks"

for i in "${FLATPAKS[@]}"; do
    flatpak install flathub $i
done
