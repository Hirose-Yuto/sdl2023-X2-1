package tools

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"main/app/models"
	"strconv"
)

type ObjectService struct {
	objectIO *ObjectIo
}

func NewObjectService(repoPath string) *ObjectService {
	return &ObjectService{objectIO: NewObjectIO(repoPath)}
}

func (oS *ObjectService) ReadBlob(objectID string) (string, error) {
	data, err := oS.objectIO.ReadObject(objectID, models.BLOB)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (oS *ObjectService) ReadTree(objectID string) (*models.TreeObject, error) {
	treeObject := &models.TreeObject{
		Elements: make([]*models.TreeElement, 0),
	}

	data, err := oS.objectIO.ReadObject(objectID, models.TREE)
	if err != nil {
		return nil, err
	}

	dataBuffer := bytes.NewBuffer(data)
	for dataBuffer.Len() > 0 {
		meta, err := oS.readStringUntilSpecificByte(dataBuffer, 32)
		if err != nil {
			return nil, err
		}
		n, err := strconv.Atoi(meta)
		if err != nil {
			return nil, err
		}
		meta = fmt.Sprintf("%06d", n)

		name, err := oS.readStringUntilSpecificByte(dataBuffer, 0)
		if err != nil {
			return nil, err
		}

		objectIDStr, err := oS.readRawHex(dataBuffer)
		if err != nil {
			return nil, err
		}

		treeObject.Elements = append(treeObject.Elements, &models.TreeElement{Meta: meta, Name: name, ObjectID: objectIDStr})
	}
	return treeObject, nil
}

func (oS *ObjectService) ReadCommit(objectID string) (*models.CommitObject, error) {
	commit := &models.CommitObject{
		Tree:      "",
		Parent:    "",
		Author:    "",
		Committer: "",
		Message:   "",
	}

	data, err := oS.objectIO.ReadObject(objectID, models.COMMIT)
	if err != nil {
		return nil, err
	}
	dataBuffer := bytes.NewBuffer(data)

	for {
		if readByte, err := dataBuffer.ReadByte(); err != nil {
			return nil, err
		} else if readByte == 10 {
			break
		}
		if err := dataBuffer.UnreadByte(); err != nil {
			return nil, err
		}

		index, err := oS.readStringUntilSpecificByte(dataBuffer, 32)
		if err != nil {
			return nil, err
		}
		switch index {
		case "tree":
			tree, err := oS.readHex(dataBuffer)
			if err != nil {
				return nil, err
			}
			commit.Tree = tree
			// 改行読み飛ばし
			if _, err := dataBuffer.ReadByte(); err != nil {
				return nil, err
			}
			break
		case "parent":
			parent, err := oS.readHex(dataBuffer)
			if err != nil {
				return nil, err
			}
			commit.Parent = parent
			// 改行読み飛ばし
			if _, err := dataBuffer.ReadByte(); err != nil {
				return nil, err
			}
			break
		case "author":
			author, err := oS.readStringUntilSpecificByte(dataBuffer, 10)
			if err != nil {
				return nil, err
			}
			commit.Author = author
			break
		case "committer":
			committer, err := oS.readStringUntilSpecificByte(dataBuffer, 10)
			if err != nil {
				return nil, err
			}
			commit.Committer = committer
			break
		}
	}

	commit.Message = dataBuffer.String()

	return commit, nil
}

func (oS *ObjectService) readStringUntilSpecificByte(dataBuffer *bytes.Buffer, sep byte) (string, error) {
	content := ""
	for {
		b, err := dataBuffer.ReadByte()
		if err != nil {
			return "", err
		}
		if b == sep {
			break
		}
		content += string(b)
	}
	return content, nil
}

func (oS *ObjectService) readHex(dataBuffer *bytes.Buffer) (string, error) {
	h := ""
	for i := 0; i < 160/4; i++ {
		b, err := dataBuffer.ReadByte()
		if err != nil {
			return "", err
		}
		h += string(b)
	}
	return h, nil
}

func (oS *ObjectService) readRawHex(dataBuffer *bytes.Buffer) (string, error) {
	h := ""
	for i := 0; i < 160/8; i++ {
		b, err := dataBuffer.ReadByte()
		if err != nil {
			return "", err
		}
		h += fmt.Sprintf("%02s", strconv.FormatInt(int64(b), 16))
	}
	return h, nil
}

func (oS *ObjectService) WriteBlob(content string) (string, error) {
	return oS.objectIO.WriteObject(content, models.BLOB)
}

func (oS *ObjectService) WriteTree(tree *models.TreeObject) (string, error) {
	content := ""
	for _, element := range tree.Elements {
		oID, err := hex.DecodeString(element.ObjectID)
		if err != nil {
			return "", err
		}
		content += element.Meta + " " + element.Name + "\000" + string(oID)
	}
	return oS.objectIO.WriteObject(content, models.TREE)
}

func (oS *ObjectService) WriteCommit(commit *models.CommitObject) (string, error) {
	content := "tree " + commit.Tree + "\n"
	if commit.Parent != "" {
		content += "parent " + commit.Parent + "\n"
	}
	content += "author " + commit.Author + "\n"
	content += "committer " + commit.Committer + "\n"
	content += "\n"
	content += commit.Message + "\n"
	return oS.objectIO.WriteObject(content, models.COMMIT)
}

func (oS *ObjectService) UpdateRef(branchName string, commitId string) error {
	return oS.objectIO.UpdateRef(branchName, commitId)
}
