package episode

import (
	"testing"
)

func TestFindEpisode(t *testing.T) {
	sampleURL := "https://freefamilyguy.com/wp-content/uploads/2023/02/Family-Guy-S03E10-Fish-Out-Of-Water.mp4"
	mockDB := Database{
		"season :22": []Episode{
			{
				Episode:        2,
				EpisodeName:    "Testing Episode",
				EpisodeMP4Link: &sampleURL,
			},
			{
				Episode:        3,
				EpisodeName:    "No MP4 Episode",
				EpisodeMP4Link: nil, 
			},
		},
	}

	tests := []struct {
		name        string
		season      string
		episode     string
		wantErr     bool
		expectedURL string
	}{
		{
			name:        "successful case: searching for an existing episode",
			season:      "s22",
			episode:     "e2",
			wantErr:     false,
			expectedURL: sampleURL,
		},
		{
			name:        "wrong season number",
			season:      "s99",
			episode:     "e1",
			wantErr:     true,
			expectedURL: "",
		},
		{
			name:        "wrong episode number",
			season:      "s22",
			episode:     "e3",
			wantErr:     true,
			expectedURL: "",
		},
		{
            name:        "capital letters case: S22 and E2",
            season:      "S22",
            episode:     "E2",
            wantErr:     false,
            expectedURL: sampleURL,
        },
		{
			name:        "leading zeros case: s022 and e002",
			season:      "s22",
			episode:     "e02",
			wantErr:     false,
			expectedURL: sampleURL,
		},
		{
			name:        "invalid format case: s22e2",
			season:      "s22e2",
			episode:     "",
			wantErr:     true,
			expectedURL: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, _, err := FindEpisode(mockDB, tt.season, tt.episode)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindEpisode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if url != tt.expectedURL {
				t.Errorf("FindEpisode() gotURL = %v, want %v", url, tt.expectedURL)
			}
		})
	}
}