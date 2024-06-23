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

func (c *command) setArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("project name is required")
	}
	c.projectName = args[0]

	return nil
}

func (c *command) run(cmd *cobra.Command, args []string) {
	cloneErr := c.cloneProject()
	c.errorHandling(cloneErr, "cloning project")

	gitFolderErr := c.removeGitFolder()
	c.errorHandling(gitFolderErr, "Removing .git directory")

	gitInitErr := c.initClearGitFolder()
	c.errorHandling(gitInitErr, "Initializing git repository")
}

func (c *command) errorHandling(err error, context string) {
	if err != nil {
		fmt.Println("Error on " + context)
		panic(err)
	}
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

func (c *command) cloneProject() error {
	fmt.Println("Creating Project in Folder" + c.projectName)
	_, err := git.PlainClone("./"+c.projectName, false, &git.CloneOptions{
		URL: "https://github.com/not-empty/zord-microframework-golang",
	})

	if err != nil {
		return err
	}
	fmt.Println("Created")

	return nil
}

func (c *command) removeGitFolder() error {
	fmt.Println("Creating regenerating .git folder without refs")
	return os.RemoveAll("./" + c.projectName + "/.git")
}

func (c *command) initClearGitFolder() error {
	_, err := git.PlainInit("./"+c.projectName+"/.git", false)
	return err
}
