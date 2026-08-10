package main


type Episode struct {
	Episode        int     `json:"Episode"`
	EpisodeName    string  `json:"Episode_Name"`
	EpisodeMP4Link *string `json:"Episode_MP4_link"` 
	
}

type Database map[string][]Episode