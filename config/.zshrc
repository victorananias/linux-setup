export ZSH="/home/victor/.oh-my-zsh"

ZSH_THEME="robbyrussell"

plugins=(
    git
    zsh-syntax-highlighting
    zsh-nvm
)

source $ZSH/oh-my-zsh.sh

export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
