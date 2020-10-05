#!/bin/bash

URL_OH_MY_ZSH="https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh"
URL_ZSH_SYNTAX_HIGHLIGHT="https://github.com/zsh-users/zsh-syntax-highlighting.git"
URL_NVM="https://github.com/lukechilds/zsh-nvm"

# install oh my zsh
sh -c "$(wget -O -fsSL $URL_OH_MY_ZSH)" "" --unattended

# download oh my zsh plugins
git clone $URL_NVM ~/.oh-my-zsh/custom/plugins/zsh-nvm
git clone $URL_ZSH_SYNTAX_HIGHLIGHT ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# set zsh as default shell
chsh -s $(which zsh)