#!/bin/bash

# set zsh as default shell
chsh -s $(which zsh)

mkdir ~/zsh
mkdir ~/zsh/extensions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ~/zsh/extensions/zsh-syntax-highlighting
