package main

const rawJsonURL = "https://gist.githubusercontent.com/eltaweel068/2b34409256da41f3e3cbc4419639937d/raw/episodes.json"

type Episode struct {
	Episode        int     `json:"Episode"`
	EpisodeName    string  `json:"Episode_Name"`
	EpisodeMP4Link *string `json:"Episode_MP4_link"` 
	
}

type Database map[string][]Episode