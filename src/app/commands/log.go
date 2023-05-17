package commands

import "fmt"

func (app *App) Log() error {
	return app.treeLog(app.ref)
}

func (app *App) treeLog(commitId string) error {
	commit, err := app.objectService.ReadCommit(commitId)
	if err != nil {
		return err
	}

	fmt.Println()

	fmt.Printf("tree: %s\n", commit.Tree)
	fmt.Printf("parent: %s\n", commit.Parent)
	fmt.Printf("author: %s\n", commit.Author)
	fmt.Printf("committer: %s\n\n", commit.Committer)
	fmt.Printf("%s\n", commit.Message)

	fmt.Println("========================")

	if commit.Parent != "" {
		return app.treeLog(commit.Parent)
	}
	return nil
}
