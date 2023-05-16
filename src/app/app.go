package app

import (
	"main/app/tools"
	"os"
	"path"
)

const repoPath = "/exec/.git"

type App struct {
	branchName  string
	ref         string
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
	ref := string(bs)
	objectService := tools.NewObjectService(repoPath)
	commit, err := objectService.ReadCommit(ref)
	if err != nil {
		return nil, err
	}

	return &App{
		branchName:  branchName,
		ref:         ref,
		currentTree: commit.Tree,

		objectService: objectService,

		stagedTree:      commit.Tree,
		pathObjectIdMap: map[string]string{},
	}, nil
}
