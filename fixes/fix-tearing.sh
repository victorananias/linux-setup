#!/bin/bash
sudo touch /etc/modprobe.d/nvidia-drm-nomodeset.conf
echo 'options nvidia-drm modeset=1' | sudo tee -a /etc/modprobe.d/nvidia-drm-nomodeset.conf
