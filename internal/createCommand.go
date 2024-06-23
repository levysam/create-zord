package internal

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func CreateCliCommand() *cobra.Command {
	cmd := &command{}
	return &cobra.Command{
		Use:   "create-project",
		Short: "Create a new zord project",
		Long:  "Create a new zord project",
		Args:  cmd.setArgs,
		Run:   cmd.run,
	}
}

type command struct {
	projectName string
}

func (c *command) run(cmd *cobra.Command, args []string) {
	fmt.Println("Creating Project in Folder" + c.projectName)
	_, err := git.PlainClone("./"+c.projectName, false, &git.CloneOptions{
		URL: "https://github.com/not-empty/zord-microframework-golang",
	})
	fmt.Println("Created")
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating regenerating .git folder without refs")
	err = os.RemoveAll("./" + c.projectName + "/.git")
	if err != nil {
		panic(err)
	}

	_, err = git.PlainInit("./"+c.projectName+"/.git", false)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}

func (c *command) setArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("project name is required")
	}
	c.projectName = args[0]

	return nil
}

func (c *command) runCommandOnProjectFolder(command string, args ...string) error {
	clicmd := exec.Command(command, args...)
	clicmd.Path = "./" + c.projectName
	err := clicmd.Run()
	if err != nil {
		return err
	}
	return nil
}
