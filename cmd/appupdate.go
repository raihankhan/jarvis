package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	DiscordURL = "https://discord.com/api/download?platform=linux&format=deb"
	ChromeURL  = "https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb"
)

func NewAppUpdateCommand() *cobra.Command {
	var appUpdateCmd = &cobra.Command{
		Use:   "appupdate",
		Short: "Update applications like Discord or Chrome.",
	}

	appUpdateCmd.AddCommand(newDiscordUpdateCommand())
	appUpdateCmd.AddCommand(newChromeUpdateCommand())

	return appUpdateCmd
}

func newDiscordUpdateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "discord",
		Short: "Download and install the latest version of Discord.",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Updating Discord...")
			downloadAndInstall(DiscordURL, "discord.deb")
		},
	}
}

func newChromeUpdateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "chrome",
		Short: "Download and install the latest version of Chrome.",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Updating Chrome...")
			downloadAndInstall(ChromeURL, "chrome.deb")
		},
	}
}

func downloadAndInstall(url, fileName string) {
	destination := filepath.Join(os.TempDir(), fileName)
	if err := downloadFile(url, destination); err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}
	defer os.Remove(destination)

	if err := installPackage(destination); err != nil {
		log.Fatalf("Failed to install package: %v", err)
	}

	log.Println("Installation complete.")
}

func downloadFile(url, destination string) error {
	cmd := exec.Command("wget", "-O", destination, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func installPackage(filePath string) error {
	cmd := exec.Command("sudo", "dpkg", "-i", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
