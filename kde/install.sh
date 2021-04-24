#!/bin/bash

declare -a APT_PACKAGES=("kio-gdrive" "latte-dock")

# install apt packages
for i in "${APT_PACKAGES[@]}"
do
    sudo apt install -y $i
done

