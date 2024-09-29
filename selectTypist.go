package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) selectTypist(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				m.mobLeader = m.members[m.cursor]
			}
		case "down", "j":
			if m.cursor < len(m.members)-1 {
				m.cursor++
				m.mobLeader = m.members[m.cursor]
			}
		case "enter":
			m.step++
		}
	}
	return m, nil
}

func (m model) selectTypistView() string {
	s := "\nselect the typist!\n\n"

	// Iterate over our choices
	for i, nowMember := range m.members {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if nowMember == m.typist {
			checked = "*" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, nowMember.name)
	}

	// The footer
	s += "\nPress q to quit.\n"
	return s
}
