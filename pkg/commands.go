package pkg

import (
	"github.com/spf13/cobra"
)

type CliCommands struct {
	cobra *cobra.Command
}

func NewCli() *CliCommands {
	return &CliCommands{
		cobra: &cobra.Command{},
	}
}

func (p *CliCommands) AddCommand(cmd *cobra.Command) {
	p.cobra.AddCommand(cmd)
}

func (p *CliCommands) Execute() error {
	err := p.cobra.Execute()
	if err != nil {
		return err
	}
	return nil
}
