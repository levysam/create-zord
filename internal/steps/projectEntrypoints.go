package steps

import (
	"create-zord/internal/ui"
	"errors"
	tea "github.com/charmbracelet/bubbletea"
)

var options = []ui.Item{
	{Title: "http", Desc: "Http implementation entrypoint", Flag: "https://github.com/levysam/zord-lambda-cmd"},
	{Title: "lambda", Desc: "Lambda implementation entrypoint (needs http implementation)", Flag: ""},
}

func GetCmdOptions() (map[string]bool, error) {
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
