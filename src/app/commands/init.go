package commands

import (
	"main/app/models"
	"strconv"
	"time"
)

func (app *App) Init() error {
	tree, err := app.objectService.WriteTree(&models.TreeObject{Elements: make([]*models.TreeElement, 0)})
	if err != nil {
		return err
	}
	commitId, err := app.objectService.WriteCommit(&models.CommitObject{
		Tree:      tree,
		Parent:    "",
		Author:    "Go-Git <go@go.com> " + strconv.Itoa(int(time.Now().Unix())) + " +0900",
		Committer: "Go-Git <go@go.com> " + strconv.Itoa(int(time.Now().Unix())) + " +0900",
		Message:   "initial auto commit",
	})
	if err != nil {
		return err
	}
	if err := app.objectService.UpdateRef(app.branchName, commitId); err != nil {
		return err
	}
	app.ref = commitId
	app.currentTree = tree
	app.stagedTree = tree
	return nil
}
