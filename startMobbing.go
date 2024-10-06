package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type stopwatchKeymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

func (m model) startMobbing(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "s", "S":
			key.WithKeys("s")
			key.
		case "t", "T":
			key.WithKeys("t")
		case "enter":
			m.step++
		}
	}
	return m, nil
}

func (m model) startMobbingView() string {
	s := "\nselect the typist!\n\n"

	// Iterate over our choices
	for i, nowMember := range m.mobSetting.members {
		if m.mobSetting.members[m.cursor] == m.mobSetting.mobLeader {
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor++
			}
		}
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		isTypist := " " // not selected
		if nowMember == m.mobSetting.typist {
			isTypist = "*" // selected!
		}
		isMobLeader := " "
		if nowMember == m.mobSetting.mobLeader {
			isMobLeader = "x"
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, strings.TrimSpace(isTypist+isMobLeader), nowMember.name)
	}

	// The footer
	s += "\nPress q to quit.\n"
	return s
}
