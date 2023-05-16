package tools

import (
	"bytes"
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

func (oS ObjectService) ReadBlob(objectID string) (string, error) {
	data, err := oS.objectIO.ReadObject(objectID, models.BLOB)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (oS ObjectService) ReadTree(objectID string) (*models.TreeObject, error) {
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

		objectIDStr, err := oS.readHexUntilSpecificByte(dataBuffer)
		if err != nil {
			return nil, err
		}

		treeObject.Elements = append(treeObject.Elements, &models.TreeElement{Meta: meta, Name: name, ObjectID: objectIDStr})
	}
	return treeObject, nil
}

func (oS ObjectService) ReadCommit(objectID string) (*models.CommitObject, error) {
	data, err := oS.objectIO.ReadObject(objectID, models.COMMIT)
	if err != nil {
		return nil, err
	}
	dataBuffer := bytes.NewBuffer(data)

	_, err = oS.readStringUntilSpecificByte(dataBuffer, 32)
	if err != nil {
		return nil, err
	}
	tree, err := oS.readHexUntilSpecificByte(dataBuffer)
	if err != nil {
		return nil, err
	}
	if _, err := dataBuffer.ReadByte(); err != nil {
		return nil, err
	}

	_, err = oS.readStringUntilSpecificByte(dataBuffer, 32)
	if err != nil {
		return nil, err
	}
	parent, err := oS.readHexUntilSpecificByte(dataBuffer)
	if err != nil {
		return nil, err
	}
	if _, err := dataBuffer.ReadByte(); err != nil {
		return nil, err
	}

	_, err = oS.readStringUntilSpecificByte(dataBuffer, 32)
	if err != nil {
		return nil, err
	}
	author, err := oS.readStringUntilSpecificByte(dataBuffer, 10)
	if err != nil {
		return nil, err
	}

	_, err = oS.readStringUntilSpecificByte(dataBuffer, 32)
	if err != nil {
		return nil, err
	}
	committer, err := oS.readStringUntilSpecificByte(dataBuffer, 10)
	if err != nil {
		return nil, err
	}

	if _, err := dataBuffer.ReadByte(); err != nil {
		return nil, err
	}

	message := dataBuffer.String()

	return &models.CommitObject{Tree: tree, Parent: parent, Author: author, Committer: committer, Message: message}, nil
}

func (oS ObjectService) readStringUntilSpecificByte(dataBuffer *bytes.Buffer, sep byte) (string, error) {
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

func (oS ObjectService) readHexUntilSpecificByte(dataBuffer *bytes.Buffer) (string, error) {
	hex := ""
	for i := 0; i < 160/8; i++ {
		b, err := dataBuffer.ReadByte()
		if err != nil {
			return "", err
		}
		hex += strconv.FormatInt(int64(b), 16)
	}
	return hex, nil
}
