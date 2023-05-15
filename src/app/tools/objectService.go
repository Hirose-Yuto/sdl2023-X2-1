package tools

import (
	"os"
	"path"
)

type ObjectService struct {
	repoPath string
}

func NewObjectService(repoPath string) *ObjectService {
	return &ObjectService{repoPath: repoPath}
}

func (oS *ObjectService) ReadBlob(objectID string) ([]byte, error) {
	bs, err := os.ReadFile(oS.getObjectPath(objectID))
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (oS *ObjectService) getObjectPath(objectID string) string {
	return path.Join(oS.repoPath, "objects", objectID[0:2], objectID[2:])
}
