package episode

import (
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"
	
)


type Episode struct {
	Episode        int     `json:"Episode"`
	EpisodeName    string  `json:"Episode_Name"`
	EpisodeMP4Link *string `json:"Episode_MP4_link"` 
	
}

type Database map[string][]Episode


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

