package main

//
//import (
//	"fmt"
//	"github.com/charmbracelet/bubbles/cursor"
//	"github.com/charmbracelet/bubbles/textarea"
//	"github.com/charmbracelet/bubbles/viewport"
//	tea "github.com/charmbracelet/bubbletea"
//	"github.com/charmbracelet/lipgloss"
//	"os"
//	"strings"
//)
//
//func chat() {
//	p := tea.NewProgram(initialModel())
//	if _, err := p.Run(); err != nil {
//		fmt.Fprintf(os.Stderr, "error: %s\n", err)
//	}
//}
//
//type model struct {
//	viewport    viewport.Model
//	messages    []string
//	textarea    textarea.Model
//	senderStyle lipgloss.Style
//	err         error
//}
//
//func initialModel() model {
//	ta := textarea.New()
//	ta.Placeholder = "Send a message..."
//	ta.Focus()
//
//	ta.Prompt = " "
//	ta.CharLimit = 280
//
//	ta.SetWidth(30)
//	ta.SetHeight(3)
//
//	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
//
//	ta.ShowLineNumbers = true
//
//	vp := viewport.New(30, 5)
//	vp.SetContent(`Welcome to the chat room! Type a message and press Enter to send`)
//
//	ta.KeyMap.InsertNewline.SetEnabled(false)
//
//	return model{
//		textarea:    ta,
//		messages:    []string{},
//		viewport:    vp,
//		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
//		err:         nil,
//	}
//}
//
//func (m model) Init() tea.Cmd {
//	return textarea.Blink
//}
//
//func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	switch msg := msg.(type) {
//	case tea.WindowSizeMsg:
//		m.viewport.Width = msg.Width
//		m.textarea.SetWidth(msg.Width)
//		return m, nil
//	case tea.KeyMsg:
//		switch msg.String() {
//		case "esc", "ctrl+c", "q":
//			//	Quit
//			fmt.Println(m.textarea.Value())
//			return m, tea.Quit
//		case "enter":
//			v := m.textarea.Value()
//			fmt.Println("textarea value:", v)
//
//			if v == "" {
//				return m, nil
//			}
//
//			m.messages = append(m.messages, m.senderStyle.Render("You: ")+v)
//			m.viewport.SetContent(strings.Join(m.messages, "\n"))
//			m.textarea.Reset()
//			m.viewport.GotoBottom()
//			return m, nil
//		default:
//			var cmd tea.Cmd
//			m.textarea, cmd = m.textarea.Update(msg)
//			return m, cmd
//		}
//	case cursor.BlinkMsg:
//		var cmd tea.Cmd
//		m.textarea, cmd = m.textarea.Update(msg)
//		return m, cmd
//
//	default:
//		return m, nil
//	}
//}
//
//func (m model) View() string {
//	return fmt.Sprintf(
//		"%s\n\n%s",
//		m.viewport.View(),
//		m.textarea.View(),
//	) + "\n\n"
//}
