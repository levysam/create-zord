package steps

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/levysam/create-zord/internal/ui"
)

var options = []ui.Item{
	{Title: "http", Desc: "Http implementation entrypoint", Flag: "http"},
	{Title: "lambda", Desc: "Lambda implementation entrypoint (needs http implementation)", Flag: "https://github.com/levysam/zord-lambda-cmd"},
}

func GetCmdOptions() (map[string]ui.Choices, error) {
	model := ui.NewMultiSelectModel("Project name:", options)
	_, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}
	if model.Exit {
		return nil, errors.New("creation canceled by user")
	}
	return model.Choices.Choices, nil
}
