#!/bin/bash

DOWNLOADS="Installation Files"
URL_GOOGLE_CHROME="https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb"
URL_OH_MY_ZSH="https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh"
URL_ZSH_SYNTAX_HIGHLIGHT="https://github.com/zsh-users/zsh-syntax-highlighting.git"
URL_NVM="https://github.com/lukechilds/zsh-nvm"
URL_UNITY_HUB="https://public-cdn.cloud.unity3d.com/hub/prod/UnityHub.AppImage"

# before execute
# move to installation directory
cd /home/$USER/Downloads/$DOWNLOADS

## add yarn repository to sourcelist
## reference https://classic.yarnpkg.com/en/docs/install/#debian-stable
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list

# update
sudo apt update -y
sudo apt upgrade -y

# install programs
sudo apt install snapd -y
sudo apt install zsh -y
sudo apt install git -y
sudo apt install fonts-firacode -y
sudo apt install vlc -y
sudo apt install steam -y
sudo apt install --no-install-recommends yarn -y
sudo apt install remmina remmina-plugin-rdp libfreerdp-plugins-standard -y

# install snaps
sudo snap install spotify 
sudo snap install code --classic
sudo snap install insomnia
sudo snap install dotnet-runtime-31
sudo snap install dotnet-sdk --classic
sudo snap install rider --classic
sudo snap alias dotnet-sdk.dotnet dotnet

# install oh my zsh
sh -c "$(wget -O -fsSL $URL_OH_MY_ZSH)" "" --unattended

# download oh my zsh plugins
git clone $URL_NVM ~/.oh-my-zsh/custom/plugins/zsh-nvm
git clone $URL_ZSH_SYNTAX_HIGHLIGHT ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

mkdir "$DOWNLOADS"
wget -c "$URL_GOOGLE_CHROME" -P "$DOWNLOADS"
wget -c "$URL_UNITY_HUB" -P "$DOWNLOADS"

# set zsh as default shell
chsh -s $(which zsh)

# install vscode sync-settings
code --install-extension Shan.code-settings-sync