package tools

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"main/app/models"
	"os"
	"path"
	"strconv"
)

type ObjectIo struct {
	repoPath string
}

func NewObjectIO(repoPath string) *ObjectIo {
	return &ObjectIo{repoPath: repoPath}
}

func (oio *ObjectIo) ReadObject(objectID string, objectType models.ObjectType) ([]byte, error) {
	bs, err := os.ReadFile(oio.getObjectPath(objectID))
	if err != nil {
		return nil, err
	}
	data, err := oio.zlibDecode(bs)
	if err != nil {
		return nil, err
	}
	data = data[len([]byte(objectType.ToString()+" ")):]
	sizeStr := ""
	dataBuffer := bytes.NewBuffer(data)
	for {
		d, err := dataBuffer.ReadByte()
		if err != nil {
			return nil, err
		}
		if d == 0 {
			break
		}
		sizeStr += string(d)
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil, err
	}
	content := make([]byte, size)
	if _, err := dataBuffer.Read(content); err != nil {
		return nil, err
	}
	return content, nil
}

func (oio *ObjectIo) WriteObject(content string, objectType models.ObjectType) (string, error) {
	blobData := objectType.ToString() + " " + strconv.Itoa(len([]byte(content))) + "\000" + content
	sh := sha1.New()
	if _, err := io.WriteString(sh, blobData); err != nil {
		return "", err
	}
	objectID := hex.EncodeToString(sh.Sum(nil))

	objectPath := oio.getObjectPath(objectID)
	if err := os.MkdirAll(path.Dir(objectPath), os.ModePerm); err != nil && !os.IsExist(err) {
		return "", err
	}

	zlibData, err := oio.zlibEncode([]byte(blobData))
	if err != nil {
		return "", nil
	}
	if err := os.WriteFile(objectPath, zlibData, os.ModePerm); err != nil {
		return "", err
	}
	return objectID, nil
}

func (oio *ObjectIo) zlibDecode(bs []byte) ([]byte, error) {
	extract, err := zlib.NewReader(bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(extract)
	return buf.Bytes(), nil
}

func (oio *ObjectIo) zlibEncode(content []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := w.Write(content); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (oio *ObjectIo) getObjectPath(objectID string) string {
	return path.Join(oio.repoPath, "objects", objectID[0:2], objectID[2:])
}

func (oio *ObjectIo) UpdateRef(branchName string, commitId string) error {
	commitId += "\n"
	if err := os.WriteFile(path.Join(oio.repoPath, branchName), []byte(commitId), os.ModePerm); err != nil {
		return err
	}
	return nil
}
