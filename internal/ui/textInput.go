package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textInput textinput.Model
	Output    string
	header    string
	Exit      bool
}

func NewInputModel(header string) *Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return &Model{
		textInput: ti,
		header:    header,
		Exit:      false,
	}
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if len(m.textInput.Value()) > 1 {
				m.Output = m.textInput.Value()
				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Exit = true
			return m, tea.Quit
		}

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n",
		m.header,
		m.textInput.View(),
	)
}
