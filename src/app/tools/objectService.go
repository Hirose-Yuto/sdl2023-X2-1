package tools

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path"
	"strconv"
)

type ObjectService struct {
	repoPath string
}

func NewObjectService(repoPath string) *ObjectService {
	return &ObjectService{repoPath: repoPath}
}

func (oS *ObjectService) ReadBlob(objectID string) (string, error) {
	bs, err := os.ReadFile(oS.getObjectPath(objectID))
	if err != nil {
		return "", err
	}
	return oS.zlibBytesToString(bs)
}

func (oS *ObjectService) WriteBlob(content string) (string, error) {
	blobData := "blob " + strconv.Itoa(len([]byte(content))) + "\000" + content
	sh := sha1.New()
	if _, err := io.WriteString(sh, blobData); err != nil {
		return "", err
	}
	objectID := hex.EncodeToString(sh.Sum(nil))

	objectPath := oS.getObjectPath(objectID)
	if err := os.MkdirAll(path.Dir(objectPath), os.ModePerm); err != nil && !os.IsExist(err) {
		return "", err
	}

	zlibData, err := oS.stringToZlibBytes(blobData)
	if err != nil {
		return "", nil
	}
	if err := os.WriteFile(objectPath, zlibData, os.ModePerm); err != nil {
		return "", err
	}
	return objectID, nil
}

func (oS *ObjectService) zlibBytesToString(bs []byte) (string, error) {
	extract, err := zlib.NewReader(bytes.NewReader(bs))
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(extract)
	return string(buf.Bytes()), nil
}

func (oS *ObjectService) stringToZlibBytes(content string) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := w.Write([]byte(content)); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (oS *ObjectService) getObjectPath(objectID string) string {
	return path.Join(oS.repoPath, "objects", objectID[0:2], objectID[2:])
}
