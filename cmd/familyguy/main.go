package main

import (
	"fmt"
	"os"
	// "os/exec"
	// "runtime"
	"log"


	"github.com/manifoldco/promptui"

	"github.com/eltaweel068/familyguy/internal/episode"
	"github.com/eltaweel068/familyguy/internal/player"
	"github.com/eltaweel068/familyguy/internal/store"
)

func main() {

	logo := `

███████╗░█████╗░███╗░░░███╗██╗██╗░░░░░██╗░░░██╗░██████╗░██╗░░░██╗██╗░░░██╗
██╔════╝██╔══██╗████╗░████║██║██║░░░░░╚██╗░██╔╝██╔════╝░██║░░░██║╚██╗░██╔╝
█████╗░░███████║██╔████╔██║██║██║░░░░░░╚████╔╝░██║░░██╗░██║░░░██║░╚████╔╝░
██╔══╝░░██╔══██║██║╚██╔╝██║██║██║░░░░░░░╚██╔╝░░██║░░╚██╗██║░░░██║░░╚██╔╝░░
██║░░░░░██║░░██║██║░╚═╝░██║██║███████╗░░░██║░░░╚██████╔╝╚██████╔╝░░░██║░░░
╚═╝░░░░░╚═╝░░╚═╝╚═╝░░░░░╚═╝╚═╝╚══════╝░░░╚═╝░░░░╚═════╝░░╚═════╝░░░░╚═╝░░░

░█████╗░██╗░░░░░██╗
██╔══██╗██║░░░░░██║
██║░░╚═╝██║░░░░░██║
██║░░██╗██║░░░░░██║
╚█████╔╝███████╗██║
░╚════╝░╚══════╝╚═╝

`

	if len(os.Args) < 3 {
	fmt.Println(logo)
	fmt.Println("Usage: familyguy <season> <episode>")
	fmt.Println("Example: familyguy s1 e5   or   familyguy 1 5")
		return
	}

	db, err := store.LoadDatabase()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	seasonInput := os.Args[1]
	episodeInput := os.Args[2]

	mp4Link, episodeName, err := episode.FindEpisode(db, seasonInput, episodeInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%s\n", episodeName)

	// cmd := exec.Command("xdg-open", mp4Link)
	// cmd.Start()
	// cmd := exec.Command("vlc", mp4Link)
	// cmd.Start()
	// fmt.Printf("MP4 Link: %s\n", mp4Link)

	// fmt.Println("Database loaded successfully! Total seasons:", len(db))

	// playInVLC(mp4Link)
	// OpenURLInBrowser(mp4Link)
	// print("use vlc to play the video or use the default browser to play the video")

	ChooseThePlayer := []string{
		"Option 1: VLC / QuickTime Player for MacOS",
		"Option 2: Default Browser",
	}

	prompt := promptui.Select{
		Label: "Select the player (Use ↑/↓ Arrow Keys)",
		Items: ChooseThePlayer,
		Size: len(ChooseThePlayer),
	}

	// Run the menu (returns index, selected item string, and error)
	index, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection canceled or failed: %v\n", err)
	}

	switch index {
	case 0:
		player.PlayInVLC(mp4Link)
	case 1:
		player.OpenURLInBrowser(mp4Link)
	}

}
