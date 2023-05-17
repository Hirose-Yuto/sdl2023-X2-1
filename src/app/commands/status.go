package commands

import "fmt"

func (app *App) Status() error {
	fmt.Println(app.branchName)
	fmt.Printf("current Tree: %s\n", app.currentTree)
	fmt.Printf("staged Tree: %s\n", app.stagedTree)
	fmt.Println()
	for path, objectId := range app.pathObjectIdMap {
		fmt.Printf("%s: %s\n", path, objectId)
	}
	return nil
}
