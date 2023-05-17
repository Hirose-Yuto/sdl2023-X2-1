package commands

import (
	"errors"
	"fmt"
	"main/app/models"
	"strconv"
	"strings"
	"time"
)

func (app *App) Commit(message string) error {
	if app.currentTree == app.stagedTree {
		return errors.New("no changes added to commit")
	}
	commit := &models.CommitObject{
		Tree:      app.stagedTree,
		Parent:    app.ref,
		Author:    "Go-Git <go@go.com> " + strconv.Itoa(int(time.Now().Unix())) + " +0900",
		Committer: "Go-Git <go@go.com> " + strconv.Itoa(int(time.Now().Unix())) + " +0900",
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

	list := strings.Split(app.branchName, "/")
	fmt.Printf("[%s %s] %s\n\n", list[len(list)-1], commitId, message)
	return nil
}
