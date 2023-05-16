package app

import (
	"main/app/models"
	"os"
	"path"
)

func (app *App) Add(args []string) error {
	for _, addedFilePath := range args {
		file, err := os.Stat(addedFilePath)
		if err != nil {
			return err
		}
		newObjId := ""
		if file.IsDir() {
			newObjId, err = app.addFolder(addedFilePath)
			if err != nil {
				return err
			}
		} else {
			newObjId, err = app.addFile(addedFilePath)
			if err != nil {
				return err
			}
		}
		for {
			dir, file := path.Split(addedFilePath)

			var tree *models.TreeObject
			if dir == "" {
				tree, err = app.objectService.ReadTree(app.stagedTree)
			} else {
				tree, err = app.objectService.ReadTree(app.pathObjectIdMap[dir])
			}
			if err != nil {
				return err
			}

			tree.UpdateObjectID(file, newObjId)
			newObjId, err = app.objectService.WriteTree(tree)
			if err != nil {
				return err
			}

			if dir == "" {
				app.stagedTree = newObjId
				break
			} else {
				addedFilePath = dir
			}
		}
	}
	return nil
}

// フォルダ全部追加してTreeを作る
func (app *App) addFolder(folder string) (string, error) {
	tree := &models.TreeObject{Elements: make([]*models.TreeElement, 0)}
	children, _ := os.ReadDir(folder)
	for _, child := range children {
		// 権限は正しくない
		if child.IsDir() {
			if treeId, err := app.addFolder(path.Join(folder, child.Name())); err != nil {
				return "", err
			} else {
				tree.Elements = append(tree.Elements, &models.TreeElement{Meta: "040000", Name: child.Name(), ObjectID: treeId})
			}
		} else {
			if objectId, err := app.addFile(path.Join(folder, child.Name())); err != nil {
				return "", err
			} else {
				tree.Elements = append(tree.Elements, &models.TreeElement{Meta: "100644", Name: child.Name(), ObjectID: objectId})
			}
		}
	}
	treeObjectId, err := app.objectService.WriteTree(tree)
	if err != nil {
		return "", err
	}
	app.pathObjectIdMap[folder] = treeObjectId
	return treeObjectId, nil
}

func (app *App) addFile(file string) (string, error) {
	bs, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	objectId, err := app.objectService.WriteBlob(string(bs))
	if err != nil {
		return "", err
	}
	app.pathObjectIdMap[file] = objectId
	return objectId, nil
}
