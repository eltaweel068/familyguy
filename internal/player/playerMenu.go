package player

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type choiceItem struct {
	title  string
	action func(string)error
}

type promptModel struct {
	url      string
	cursor   int
	choices  []choiceItem
	canceled bool
}

func newPromptModel(url string) promptModel {
	return promptModel{
		url:    url,
		cursor: 0,
		choices: []choiceItem{
			{
				title:  "VLC [You must have VLC installed]",
				action: PlayInVLC,
			},
			{
				title:  "Default Web Browser",
				action: OpenURLInBrowser,
			},
		},
	}
}

func (m promptModel) Init() tea.Cmd {
	return nil
}

func (m promptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c ", "esc":
			m.canceled = true
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case "enter":
			_ = m.choices[m.cursor].action(m.url)
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m promptModel) View() string {
	if m.canceled {
		return ""
	}

	var b strings.Builder
	b.WriteString("\n? Select the player:\n\n")

	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "▸ "
		}
		b.WriteString(fmt.Sprintf("%sOption %d: %s\n", cursor, i+1, choice.title))
	}

	b.WriteString("\n[↑/↓: Navigate • Enter: Select • Esc: Cancel]\n")
	return b.String()
}

func PromptAndPlay(url string) error {
	p := tea.NewProgram(newPromptModel(url))
	_, err := p.Run()
	return err
}

