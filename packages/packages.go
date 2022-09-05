package packages

func GetAptPackages() []string {
	return []string{"curl", "zsh", "git", "fonts-firacode", "default-jre", "default-jdk", "calibre", "qbittorrent", "flatpak", "pulseeffects"}
}

func GetFlatpaks() []string {
	return []string{"com.discordapp.Discord", "org.mozilla.firefox"}
}

func GetDebUrls() []string {
	return []string{"https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb", "https://updates.insomnia.rest/downloads/ubuntu/latest"}
}

func GetAppimages() []string {
	return []string{}
}
