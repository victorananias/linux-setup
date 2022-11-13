#!/bin/bash

flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo

declare -a FLATPAKS=(
    
)

echo "Install apt packages"

for i in "${FLATPAKS[@]}"; do
    flatpak install flathub $i -y
done
