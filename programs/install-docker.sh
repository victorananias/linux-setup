#!/bin/sh

# install dependencies
sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# add apt-repository
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# install docker
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io

# add current user to docker group
sudo usermod -aG docker $USER

# install docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.28.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/bin/docker-compose

# Apply executable permissions to docker-compose binary
sudo chmod +x /usr/bin/docker-compose