package app

import (
	"errors"
	"main/app/models"
)

func (app *App) Commit(message string) error {
	if app.currentTree == app.stagedTree {
		return errors.New("no changes added to commit")
	}
	commit := &models.CommitObject{
		Tree:      app.stagedTree,
		Parent:    app.currentTree,
		Author:    "Go Git <go@go.com> 00000 +0900",
		Committer: "Go Git <go@go.com> 00000 +0900",
		Message:   message,
	}
	commitId, err := app.objectService.WriteCommit(commit)
	if err != nil {
		return err
	}
	if err := app.objectService.UpdateRef(app.branchName, commitId); err != nil {
		return err
	}
	app.currentTree = app.stagedTree
	app.ref = commitId
	return nil
}
