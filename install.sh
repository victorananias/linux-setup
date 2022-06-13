#!/bin/bash

echo "Update"

sudo apt update -y
sudo apt upgrade -y

./install/install-apt-packages.sh
./install/install-deb-packages.sh
./install/install-snap-packages.sh
./install/run-install-programs.sh
./install/download-appimages.sh
./install/run-post-install.sh

echo "PLEASE REBOOT..."
