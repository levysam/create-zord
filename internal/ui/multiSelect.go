package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	Flag, Title, Desc string
}

type Selection struct {
	Choices map[string]Choices
}

type Choices struct {
	Name    string
	Install bool
}

func (s *Selection) Update(optionName string, value bool, flag string) {
	s.Choices[flag] = Choices{
		Name:    optionName,
		Install: value,
	}
}

type MultiSelectModel struct {
	cursor   int
	options  []Item
	selected map[int]struct{}
	Choices  Selection
	header   string
	Exit     bool
}

var (
	focusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	selectedItemStyle     = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
	selectedItemDescStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170"))
	descriptionStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#40BDA3"))
)

func NewMultiSelectModel(header string, options []Item) *MultiSelectModel {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return &MultiSelectModel{
		options:  options,
		selected: make(map[int]struct{}),
		Choices: Selection{
			Choices: map[string]Choices{},
		},
		header: header,
		Exit:   false,
	}
}

func (m *MultiSelectModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *MultiSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Exit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			for selectedKey := range m.selected {
				m.Choices.Update(m.options[selectedKey].Title, true, m.options[selectedKey].Flag)
				m.cursor = selectedKey
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *MultiSelectModel) View() string {
	s := m.header + "\n\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			option.Title = selectedItemStyle.Render(option.Title)
			option.Desc = selectedItemDescStyle.Render(option.Desc)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = focusedStyle.Render("*")
		}

		title := focusedStyle.Render(option.Title)
		description := descriptionStyle.Render(option.Desc)

		s += fmt.Sprintf("%s [%s] %s\n%s\n\n", cursor, checked, title, description)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n", focusedStyle.Render("y"))
	return s
}
