#!/bin/bash

declare -a SNAP_PACKAGES=("postman" "spotify")

echo "Installing snap packages"

for i in "${SNAP_PACKAGES[@]}"
do
    sudo snap install $i
done
