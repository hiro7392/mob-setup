package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}

}

type model struct {
	viewport    viewport.Model
	messages    []string
	members     members
	mobLeader   member
	typist      member
	step        Step
	cursor      int
	senderStyle lipgloss.Style
	err         error
}

type member struct {
	name string
}
type members []member

type Step int

const (
	SELECT_MOBLEADER Step = iota
	SELECT_TYPIST
	SHUFFLE
)

var defaultMembers members = members{
	{
		name: "hiro",
	},
	{
		name: "ichi",
	},
	{
		name: "mochi",
	},
	{
		name: "taka",
	},
}

func initialModel() model {
	vp := viewport.New(30, len(defaultMembers)+1)
	vp.SetContent(`Setting up mob programming ...`)

	return model{
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		members:     defaultMembers,
		cursor:      1,
		typist:      defaultMembers[0],
		mobLeader:   defaultMembers[1],
		step:        SELECT_MOBLEADER,
		err:         nil,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch m.step {
	case SELECT_MOBLEADER:
		return m.selectMobLeader(msg)
	case SELECT_TYPIST:
		return m.selectTypist(msg)
	case SHUFFLE:
		return m.selectMobLeader(msg)
	default:
		return m.selectMobLeader(msg)
	}
}

func (m model) View() string {
	switch m.step {
	case SELECT_MOBLEADER:
		return m.selectMobLeaderView()
	case SELECT_TYPIST:
		return m.selectTypistView()
	case SHUFFLE:
		return m.selectMobLeaderView()
	default:
		return m.selectMobLeaderView()
	}
}

func (m model) Init() tea.Cmd {
	m.senderStyle.Render("")
	return textarea.Blink
}
