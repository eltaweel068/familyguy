package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"runtime"
	"os/exec"
	"net/http"
)


func LoadDatabase() (Database, error) {
	resp, err := http.Get(rawJsonURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var db Database
	err = json.NewDecoder(resp.Body).Decode(&db)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return db, nil
}

func FindEpisode(db Database, seasonInput, episodeInput string) (string, string, error) {
	seasonInput = strings.ToLower(seasonInput)
	episodeInput = strings.ToLower(episodeInput)

	seasonInput = strings.ReplaceAll(seasonInput, "s", "")
	episodeInput = strings.ReplaceAll(episodeInput, "e", "")

	seasonKey := fmt.Sprintf("season :%s", strings.TrimSpace(seasonInput))

	epNum, err := strconv.Atoi(strings.TrimSpace(episodeInput))
	if err != nil {
		return "", "", fmt.Errorf("invalid episode number: %v", err)
	}

	episodes, exists := db[seasonKey]
	if !exists {
		return "", "", fmt.Errorf("season [%s] not found", seasonKey)
	}


	for _, ep := range episodes {
    if ep.Episode == epNum {
        if ep.EpisodeMP4Link == nil {
            return "", ep.EpisodeName, fmt.Errorf("episode [%d] does not have a valid MP4 link", epNum)
        }
        return *ep.EpisodeMP4Link, ep.EpisodeName, nil
    }
}

	return "", "", fmt.Errorf("episode [%d] not found in season [%s]", epNum, seasonKey)
}

// Function to open a URL in the default browser based on the operating system

func OpenURLInBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Start()
}

// Function to play a URL in VLC based on the operating system

func playInVLC(url string) error {
	switch runtime.GOOS {

	case "linux":
		if _, err := exec.LookPath("vlc"); err == nil {
			return exec.Command("vlc", url).Start()
		}
		if _, err := exec.LookPath("flatpak"); err == nil {
			cmd := exec.Command("flatpak", "run", "org.videolan.VLC", url)
			if err := cmd.Start(); err == nil {
				return nil
			}
		}
		if _, err := exec.LookPath("snap"); err == nil {
			cmd := exec.Command("snap", "run", "vlc", url)
			if err := cmd.Start(); err == nil {
				return nil
			}
		}

	case "windows":
		if _, err := exec.LookPath("vlc"); err == nil {
			return exec.Command("vlc", url).Start()
		}


		defaultWindowsPaths := []string{
			`C:\Program Files\VideoLAN\VLC\vlc.exe`,
			`C:\Program Files (x86)\VideoLAN\VLC\vlc.exe`,
			//`C:\Program Files\VLC\vlc.exe`,
			//`C:\Program Files (x86)\VLC\vlc.exe`,
		}
		// is the file exists at the given path
		for _, path := range defaultWindowsPaths {
			if _, err := os.Stat(path); err == nil {
				return exec.Command(path, url).Start()
			}
		}

	case "darwin":
		cmd := exec.Command("open", "-a", "QuickTime Player", url)
		if err := cmd.Start(); err == nil {
			return nil
		}
		
		if _, err := exec.LookPath("vlc"); err == nil {
			return exec.Command("vlc", url).Start()
		}
	}

	return fmt.Errorf("VLC not found or unsupported operating system: %s", runtime.GOOS)
}