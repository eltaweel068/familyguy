package store

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/eltaweel068/familyguy/internal/episode"
)

const RawJsonURL = "https://gist.githubusercontent.com/eltaweel068/2b34409256da41f3e3cbc4419639937d/raw/episodes.json"

func LoadDatabase() (episode.Database, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(RawJsonURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching episodes: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-200 status: %d", resp.StatusCode)
	}

	var db episode.Database
	err = json.NewDecoder(resp.Body).Decode(&db)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return db, nil
}