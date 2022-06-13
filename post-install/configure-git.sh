#!/bin/bash

GIT_NAME="Victor"
GIT_EMAIL="victor@ananias.dev"
GITHUB_URL="https://github.com/settings/keys"

git config --global user.name "$GIT_NAME"
git config --global user.email "$GIT_EMAIL"

# creates a new ssh key
ssh-keygen -t rsa -b 4096 -C "$GIT_EMAIL"

# start the ssh-agent 
eval "$(ssh-agent -s)"

# add SSH private key to the ssh-agent
ssh-add ~/.ssh/id_rsa


# printing public ssh
cat ~/.ssh/id_rsa.pub

# https://github.com/settings/keys
echo "paste your ssh key into your git provider here: $GITHUB_URL"
