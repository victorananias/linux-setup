#!/bin/bash

echo "Update"

sudo apt update -y
sudo apt upgrade -y

./Install/apt-packages.sh
./Install/deb-packages.sh
./Install/flatpaks.sh
./Install/run-post-install.sh

echo "PLEASE REBOOT..."
