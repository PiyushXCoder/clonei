package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/soft4dev/iclone/internal/projects"
	"github.com/spf13/cobra"
)

var (
	project string
	install bool
	cd      bool
)

var rootCmd = &cobra.Command{
	Use:   "iclone",
	Short: "clone and install deps of project",
	Long: `
		It clones provided repo using git and install dependencies according to project type. eg. npm, pnpm, go, rust....
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if r := checkGitInstalled(); r != nil {
			fmt.Println("Git is not installed")
			os.Exit(1)
		}
		repoUrl := args[0]
		gitCloneOutput := exec.Command("git", "clone", repoUrl)
		gitCloneOutput.Stdout = os.Stdout
		gitCloneOutput.Stderr = os.Stderr
		gitCloneOutput.Stdin = os.Stdin
		if err := gitCloneOutput.Run(); err != nil {
			fmt.Println("Error cloning repo:", err)
			return
		}

		// Extract the directory name from the repo URL
		// e.g., "https://github.com/user/repo.git" -> "repo"
		projectDirName := repoUrl
		if idx := strings.LastIndex(projectDirName, "/"); idx != -1 {
			projectDirName = projectDirName[idx+1:]
		}
		projectDirName = strings.TrimSuffix(projectDirName, ".git")
		projectType := ""
		if project == "AUTO" {
			var err error
			if projectType, err = projects.ProjectDetector(projectDirName); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		} else {
			projectType = strings.ToUpper(project)
		}
		handler := projects.ProjectHandlers[projectType]
		if handler == nil {
			fmt.Printf("Error: No handler found for project type '%s'\n", projectType)
			fmt.Println("Available project types:", projects.GetAvailableProjectTypes())
			os.Exit(1)
		}
		if err := handler.Install(projectDirName); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if cd {
			if err := os.Chdir(projectDirName); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}

		fmt.Println("project: " + projectType)
		fmt.Println("url: " + args[0])

	},
	Args: cobra.ExactArgs(1),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&project, "project", "p", "AUTO", "Project type (npm, go, rust, etc.). Use AUTO for auto-detection")
	rootCmd.Flags().BoolVarP(&install, "install", "i", true, "controls whether to install dependencies after clone")
	rootCmd.Flags().BoolVarP(&cd, "cd", "c", true, "controls whether to change directory into the project folder after clone")
}

func checkGitInstalled() error {
	cmd := exec.Command("git", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git is not installed or not available in PATH")
	}
	return nil
}
