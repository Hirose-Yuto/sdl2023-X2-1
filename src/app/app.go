package app

import (
	"main/app/tools"
	"os"
	"path"
)

const repoPath = "/exec/.git"

type App struct {
	branchName  string
	currentTree string

	// Services
	objectService *tools.ObjectService

	// Internal Data
	stagedTree      string
	pathObjectIdMap map[string]string
}

func NewApp() (*App, error) {
	branchName := "refs/heads/main"
	bs, err := os.ReadFile(path.Join(repoPath, branchName))
	if err != nil {
		return nil, err
	}
	return &App{
		branchName:  branchName,
		currentTree: string(bs),

		objectService: tools.NewObjectService(repoPath),

		stagedTree:      string(bs),
		pathObjectIdMap: map[string]string{},
	}, nil
}
