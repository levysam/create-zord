package internal

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/levysam/create-zord/internal/steps"
	"github.com/levysam/create-zord/internal/ui"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
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
	ProjectEntryPoints map[string]ui.Choices
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

	replaceErr := c.replaceProjectName("./" + c.ProjectName)
	c.errorHandling(replaceErr, "replacing project")

	gitFolderErr := c.removeInProjectFolder("/.git")
	c.errorHandling(gitFolderErr, "Removing .git directory")

	gitInitErr := c.initClearGitFolder()
	c.errorHandling(gitInitErr, "Initializing git repository")
}

func (c *Command) errorHandling(err error, context string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", context, err)
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
	for key, choice := range c.ProjectEntryPoints {
		err := c.installEntrypoint(key, choice.Install, choice.Name)
		c.errorHandling(err, "installing entrypoint")
	}
	return nil
}

func (c *Command) installEntrypoint(url string, install bool, name string) error {
	if url == "http" {
		if !install {
			return c.removeInProjectFolder("/cmd/http")
		}
		return nil
	}

	err := c.cloneProject("./"+c.ProjectName+"/cmd/"+name, url)
	if err != nil {
		return err
	}
	return nil
}

func (c *Command) replaceProjectName(projectPath string) error {
	return filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return c.replaceNameInFile(path)
	})
}

func (c *Command) replaceNameInFile(path string) error {
	data, err := c.getFileData(path)
	if err != nil {
		return err
	}
	replacedString := strings.ReplaceAll(data, "zord", c.ProjectName)
	return os.WriteFile(path, []byte(replacedString), 0644)
}

func (c *Command) getFileData(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
