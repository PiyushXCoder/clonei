package projects

import (
	"errors"
	"os"
	"os/exec"
)

func GetAvailableProjectTypes() []string {
	types := make([]string, 0, len(ProjectHandlers))
	for k := range ProjectHandlers {
		types = append(types, k)
	}
	return types
}

type ProjectHandler interface {
	Install(projectDir string) error
}

var ProjectHandlers = map[string]ProjectHandler{
	"NPM":  npmHandler{},
	"PNPM": pnpmHandler{},
}

type npmHandler struct{}
type pnpmHandler struct{}

func (n npmHandler) Install(projectDir string) error {
	// TODO: check npm command
	init := exec.Command("npm", "ci")
	init.Dir = projectDir
	init.Stdout = os.Stdout
	init.Stderr = os.Stderr
	init.Stdin = os.Stdin
	if err := init.Run(); err != nil {
		return errors.New("error initializing project")
	}

	return nil
}
func (n pnpmHandler) Install(projectDir string) error {
	// TODO: check pnpm command
	init := exec.Command("pnpm", "install", "--frozen-lockfile")
	init.Dir = projectDir
	init.Stdout = os.Stdout
	init.Stderr = os.Stderr
	init.Stdin = os.Stdin
	if err := init.Run(); err != nil {
		return errors.New("error initializing project")
	}

	return nil
}
