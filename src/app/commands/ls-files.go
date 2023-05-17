package commands

import (
	"fmt"
	"path"
)

func (app *App) LsFiles() error {
	list, err := app.getFileList(app.currentTree, "")
	if err != nil {
		return err
	}
	for _, f := range list {
		fmt.Println(f)
	}
	return nil
}

func (app *App) getFileList(treeId string, pathName string) ([]string, error) {
	var l []string
	tree, err := app.objectService.ReadTree(treeId)
	if err != nil {
		return nil, err
	}

	for _, element := range tree.Elements {
		switch element.Meta[:2] {
		case "10":
			// ファイル
			l = append(l, path.Join(pathName, element.Name))
			break
		case "04":
			// ディレクトリ
			if l2, err := app.getFileList(element.ObjectID, path.Join(pathName, element.Name)); err != nil {
				return nil, err
			} else {
				l = append(l, l2...)
			}
			break
		}
	}
	return l, err
}
