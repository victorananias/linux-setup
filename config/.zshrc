export ZSH="$HOME/.oh-my-zsh"
export ZSH_CONFIG="$HOME/zsh-config"

ZSH_THEME="spaceship"

plugins=(
    git
    bundler
    dotenv
    macos
    rake
    rbenv
    jsontools
    node
    pip
    web-search
    zsh-autosuggestions
    colored-man-pages
    colorize
    common-aliases
    copyfile
    zsh-syntax-highlighting
)

source $ZSH/oh-my-zsh.sh

export GOROOT=$HOME/.go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT/bin:$GOBIN

# KUBECTL ALIASES
alias k="kubectl"
alias klogs="kubectl logs"
alias kpods="kubectl get pods"
alias kdpod="kubectl delete pod"
alias keditcmap=kubectl edit configmap

szs() {
    source $HOME/.zshrc
}

kill-go() {
    for each in $(ps -Al | grep -w go | awk '{print $2}'); do
        kill $each
    done
}

set-var-in-file() {
    variable="${1}"
    content="${2}"
    file="${3}"

    if [ ! -f "${file}" ]; then
        echo "set-var-in-file: file doesn't exist: ${file}"
        exit 1
    fi

    sed -i "" "s/${variable}=.*/${variable}=\"${content}\"/" "$file"
}

for each in $(ls $ZSH_CONFIG); do
    source "$ZSH_CONFIG/$each"
done
