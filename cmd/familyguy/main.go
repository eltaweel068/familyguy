package main

import (
	"fmt"
	"os"

	// "os/exec"
	// "runtime"
	"log"

	// "github.com/manifoldco/promptui"

	"github.com/eltaweel068/familyguy/internal/episode"
	"github.com/eltaweel068/familyguy/internal/player"
	"github.com/eltaweel068/familyguy/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eltaweel068/familyguy/internal/ui"
)



func main() {

	// if len(os.Args) < 3 {
	// fmt.Print(Logo)
	// fmt.Println("Usage: familyguy <season> <episode>")
	// fmt.Println("Example: familyguy s1 e5   or   familyguy 1 5")
	// 	return
	// }

	db, err := store.LoadDatabase()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}


	var targetLink string

	if len(os.Args) >= 3 {
		seasonInput := os.Args[1]
		episodeInput := os.Args[2]

		link, episodeName, err := episode.FindEpisode(db, seasonInput, episodeInput)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Playing: %s\n", episodeName)
		targetLink = link

	} else {
		p := tea.NewProgram(ui.NewModel(db))
		finalModel, err := p.Run()
		if err != nil {
			log.Fatalf("Error running program: %v\n", err)
		}

		m, ok := finalModel.(ui.Model)
		if ok && m.SelectedLink != "" {
			targetLink = m.SelectedLink
		}
	}

	if targetLink != "" {
		if err := player.PromptAndPlay(targetLink); err != nil {
			log.Printf("Error: %v\n", err)
		}
	}


	// p := tea.NewProgram(ui.NewModel(db))
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Error running program: %v\n", err)
	// 	os.Exit(1)
	// }

	// seasonInput := os.Args[1]
	// episodeInput := os.Args[2]

	// mp4Link, episodeName, err := episode.FindEpisode(db, seasonInput, episodeInput)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("%s\n", episodeName)

	// cmd := exec.Command("xdg-open", mp4Link)
	// cmd.Start()
	// cmd := exec.Command("vlc", mp4Link)
	// cmd.Start()
	// fmt.Printf("MP4 Link: %s\n", mp4Link)

	// fmt.Println("Database loaded successfully! Total seasons:", len(db))

	// playInVLC(mp4Link)
	// OpenURLInBrowser(mp4Link)
	// print("use vlc to play the video or use the default browser to play the video")

	// Run the menu (returns index, selected item string, and error)

}







// func playerMassge(url string) {
// 		ChooseThePlayer := []string{
// 		"Option 1: VLC / QuickTime Player for MacOS",
// 		"Option 2: Default Browser",
// 	}

// 	prompt := promptui.Select{
// 		Label: "Select the player (Use ↑/↓ Arrow Keys)",
// 		Items: ChooseThePlayer,
// 		Size:  len(ChooseThePlayer),
// 	}

// 	index, _, err := prompt.Run()
// 	if err != nil {
// 		log.Fatalf("Selection canceled or failed: %v\n", err)
// 	}

// 	switch index {
// 	case 0:
// 		player.PlayInVLC(url)
// 	case 1:
// 		player.OpenURLInBrowser(url)
// 	}
// }
