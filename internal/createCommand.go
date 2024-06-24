package internal

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/levysam/create-zord/internal/steps"
	"github.com/spf13/cobra"
	"os"
)

func CreateCliCommand() *cobra.Command {
	cmd := &Command{}
	return &cobra.Command{
		Use:   "create-project",
		Short: "Create a new zord project",
		Long:  "Create a new zord project",
		Run:   cmd.run,
	}
}

type Command struct {
	ProjectName        string
	ProjectEntryPoints map[string]bool
}

func (c *Command) run(cmd *cobra.Command, args []string) {
	name, nameErr := steps.GetProjectName()
	c.errorHandling(nameErr, "Step Project Name")
	c.ProjectName = name

	entryPoints, EntryErr := steps.GetCmdOptions()
	c.errorHandling(EntryErr, "Step Cmd options")
	c.ProjectEntryPoints = entryPoints

	fmt.Println("Creating Project")

	cloneErr := c.cloneProject(c.ProjectName, "https://github.com/not-empty/zord-microframework-golang")
	c.errorHandling(cloneErr, "cloning project")

	addCmdErr := c.addZordEntrypoint()
	c.errorHandling(addCmdErr, "adding entrypoint")

	gitFolderErr := c.removeInProjectFolder("/.git")
	c.errorHandling(gitFolderErr, "Removing .git directory")

	gitInitErr := c.initClearGitFolder()
	c.errorHandling(gitInitErr, "Initializing git repository")
}

func (c *Command) errorHandling(err error, context string) {
	if err != nil {
		os.Exit(1)
	}
}

func (c *Command) cloneProject(path string, url string) error {
	_, err := git.PlainClone("./"+path, false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *Command) removeInProjectFolder(path string) error {
	return os.RemoveAll("./" + c.ProjectName + path)
}

func (c *Command) initClearGitFolder() error {
	_, err := git.PlainInit("./"+c.ProjectName+"/.git", false)
	return err
}

func (c *Command) addZordEntrypoint() error {
	for key, install := range c.ProjectEntryPoints {
		err := c.installEntrypoint(key, install)
		c.errorHandling(err, "installing entrypoint")
	}
	return nil
}

func (c *Command) installEntrypoint(url string, install bool) error {
	if url == "http" && !install {
		removeErr := c.removeInProjectFolder("/cmd/http")
		return removeErr
	}

	err := c.cloneProject("./"+c.ProjectName+"/cmd", url)
	if err != nil {
		return err
	}
	return nil
}
