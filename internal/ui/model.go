package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/eltaweel068/familyguy/internal/episode"
		// "github.com/eltaweel068/familyguy/internal/player"

	"sort"
	"strconv"
	"strings"
	
)

const Logo = `
███████╗ █████╗ ███╗   ███╗██╗██╗  ██╗   ██╗   ██████╗ ██╗   ██╗██╗   ██╗
██╔════╝██╔══██╗████╗ ████║██║██║  ╚██╗ ██╔╝  ██╔════╝ ██║   ██║╚██╗ ██╔╝
█████╗  ███████║██╔████╔██║██║██║   ╚████╔╝   ██║  ███╗██║   ██║ ╚████╔╝ 
██╔══╝  ██╔══██║██║╚██╔╝██║██║██║    ╚██╔╝    ██║   ██║██║   ██║  ╚██╔╝  
██║     ██║  ██║██║ ╚═╝ ██║██║███████╗██║     ╚██████╔╝╚██████╔╝   ██║   
╚═╝     ╚═╝  ╚═╝╚═╝     ╚═╝╚═╝╚══════╝╚═╝      ╚═════╝  ╚═════╝    ╚═╝   

`

func extractSeasonNum(s string) int {
	numStr := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(s, "season", ""), ":", ""))
	num, _ := strconv.Atoi(numStr)
	return num
}

type Model struct {
	seasons       []string
	cursor        int
	db            episode.Database
	episodeCursor int
	seasonCursor  int
	state         bool
	SelectedLink  string
}

func NewModel(db episode.Database) Model {
	var seasons []string
	for s := range db {
		seasons = append(seasons, s)
	}

	sort.Slice(seasons, func(i, j int) bool {
		return extractSeasonNum(seasons[i]) < extractSeasonNum(seasons[j])
	})

	return Model{
		seasons: seasons,
		cursor:  0,
		db:      db,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c","esc":
			return m, tea.Quit

		case "up":
			if !m.state {
				if m.cursor >= 0 {
					m.cursor--
				}
				if m.cursor < 0 {
					m.cursor = len(m.seasons) - 1
				}
			} else {
				currentSeasonKey := m.seasons[m.cursor]
				totalEpisodes := len(m.db[currentSeasonKey])

				if m.episodeCursor > 0 {
					m.episodeCursor--
				} else {
					m.episodeCursor = totalEpisodes - 1
				}
			}

		case "down":
			if !m.state {
				if m.cursor <= len(m.seasons)-1 {
					m.cursor++
				}
				if m.cursor == len(m.seasons) {
					m.cursor = 0
				}
			} else {
				currentSeasonKey := m.seasons[m.cursor]
				totalEpisodes := len(m.db[currentSeasonKey])
				if m.episodeCursor < totalEpisodes-1 {
					m.episodeCursor++
				} else {
					m.episodeCursor = 0
				}
			}

		case "right":
			m.state = true
			m.episodeCursor = 0

		case "left":
			m.state = false

		case "enter":
            if !m.state {
                m.state = true
                m.episodeCursor = 0
            } else {
                currentSeasonKey := m.seasons[m.cursor]
                selectedEpisode := m.db[currentSeasonKey][m.episodeCursor]

                if selectedEpisode.EpisodeMP4Link != nil {
                    m.SelectedLink = *selectedEpisode.EpisodeMP4Link
                    return m, tea.Quit
                }
            }


		}
	}

	return m, nil
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


func (m Model) View() string {

	var seasonsBuilder strings.Builder
	seasonsBuilder.WriteString("Seasons:\n\n")

	for i, seasonKey := range m.seasons {
		seasonNum := extractSeasonNum(seasonKey)
		if m.cursor == i {
			if !m.state {
				seasonsBuilder.WriteString(fmt.Sprintf("▸ Season %02d\n", seasonNum))
			} else {
				seasonsBuilder.WriteString(fmt.Sprintf("• Season %02d\n", seasonNum))
			}
		} else {
			seasonsBuilder.WriteString(fmt.Sprintf("  Season %02d\n", seasonNum))
		}
	}

	var episodesBuilder strings.Builder
	currentSeasonKey := m.seasons[m.cursor]
	episodes := m.db[currentSeasonKey]

	episodesBuilder.WriteString(fmt.Sprintf("Episodes (%s):\n\n", m.seasons[m.cursor]))

	for j, ep := range episodes {
		epCursor := "  "
		if m.state && m.episodeCursor == j {
			epCursor = " ▸"
		}
		cleanTitle := cleanEpisodeTitle(ep.EpisodeName)
		episodesBuilder.WriteString(fmt.Sprintf("%s%02d │ %s\n", epCursor, ep.Episode, cleanTitle))
	}

	leftCol := lipgloss.NewStyle().Width(18).Render(seasonsBuilder.String())
	rightCol := episodesBuilder.String()

	mainView := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, "│ ", rightCol)

	return fmt.Sprintf("%s\n\n%s\n\nPress Ctrl+C to quit.\n", Logo, mainView)
}

func cleanEpisodeTitle(title string) string {
	parts := strings.Split(title, "|")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(title)
}

