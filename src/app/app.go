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

	pathObjectIdMap, err := getPathObjectIdMap(objectService, commit.Tree, "")

	return &App{
		branchName:  branchName,
		ref:         ref,
		currentTree: commit.Tree,

		objectService: objectService,

		stagedTree:      commit.Tree,
		pathObjectIdMap: pathObjectIdMap,
	}, nil
}

func getPathObjectIdMap(objectService *tools.ObjectService, treeId string, pathStr string) (map[string]string, error) {
	pathObjectIdMap := map[string]string{}
	tree, err := objectService.ReadTree(treeId)
	if err != nil {
		return nil, err
	}
	for _, element := range tree.Elements {
		switch element.Meta[0:2] {
		case "10":
			// ファイル, blob
			pathObjectIdMap[path.Join(pathStr, element.Name)] = element.ObjectID
			break
		case "04":
			// ディレクトリ, tree
			pathObjectIdMap2, err := getPathObjectIdMap(objectService, element.ObjectID, path.Join(pathStr, element.Name))
			if err != nil {
				return nil, err
			}
			pathObjectIdMap = merge(pathObjectIdMap, pathObjectIdMap2)
			break
		}
	}

	return pathObjectIdMap, nil
}

func merge(m ...map[string]string) map[string]string {
	ans := make(map[string]string, 0)

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}
