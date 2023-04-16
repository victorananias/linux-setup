#!/bin/bash

EMAIL="victor@ananias.dev"

# Generating a new SSH key
ssh-keygen -t ed25519 -C $EMAIL

# Adding your SSH key to the ssh-agent
eval "$(ssh-agent -s)"

ssh-add ~/.ssh/id_ed25519

# Print SSH public key
cat ~/.ssh/id_ed25519.pub