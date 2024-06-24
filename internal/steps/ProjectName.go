package steps

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/levysam/create-zord/internal/ui"
)

func GetProjectName() (string, error) {
	model := ui.NewInputModel("Project name:")
	_, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}
	if model.Exit {
		return "", errors.New("creation canceled by user")
	}
	return model.Output, nil
}
