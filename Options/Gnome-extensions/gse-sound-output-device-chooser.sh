#!/bin/bash

cd ~/.local/share/gnome-shell/extensions/

# Remove older version
rm -rf *sound-output-device-chooser*

# Clone current version
git clone https://github.com/kgshank/gse-sound-output-device-chooser.git

# Install it
cp -r gse-sound-output-device-chooser/sound-output-device-chooser@kgshank.net .
rm -rf "gse-sound-output-device-chooser"

gnome-shell-extension-tool -e gse-sound-output-device-chooser
