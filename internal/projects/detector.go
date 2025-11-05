package projects

import (
	"errors"
	"os"
	"path/filepath"
)

func ProjectDetector(projectDirName string) (string, error) {
	pnpmLockPath := filepath.Join(projectDirName, "pnpm-lock.yaml")
	if _, err := os.Stat(pnpmLockPath); err == nil {
		return "PNPM", nil
	}

	npmLockPath := filepath.Join(projectDirName, "package-lock.json")
	if _, err := os.Stat(npmLockPath); err == nil {
		return "NPM", nil
	}

	return "", errors.New("no supported project, Just cloned")
}
