#!/bin/bash

sudo apt install software-properties-gtk -y

# install openrazer
sudo add-apt-repository ppa:openrazer/stable -y && sudo apt update -y && sudo apt install openrazer-meta -y

sudo add-apt-repository ppa:polychromatic/stable -y && sudo apt update && sudo apt install polychromatic -y