#!/bin/sh

# Setup the stable repository and the GPG key
distribution=$(. /etc/os-release;echo $ID$VERSION_ID) \
   && curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | sudo apt-key add - \
   && curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | sudo tee /etc/apt/sources.list.d/nvidia-docker.list

# Install the nvidia-docker2 package (and dependencies) after updating the package listing
sudo apt-get update
sudo apt-get install -y nvidia-docker2

# Restart the Docker daemon to complete the installation after setting the default runtime
sudo systemctl restart docker

# docker run --runtime=nvidia -p 8888:8888 -v /home/victor/Development/JupyterLab:/tf --name tensorflow tensorflow/tensorflow:latest-gpu-jupyter