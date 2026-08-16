package player
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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

func PlayInVLC(url string) error {
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