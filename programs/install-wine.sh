#!/bin/bash

# enable 32 bit architecture
sudo dpkg --add-architecture i386 

# download and add the repository key:
wget -nc https://dl.winehq.org/wine-builds/winehq.key
sudo apt-key add winehq.key

# add repository for ubuntu focal main
sudo add-apt-repository 'deb https://dl.winehq.org/wine-builds/ubuntu/ focal main'

# install stable
sudo apt install --install-recommends winehq-stable