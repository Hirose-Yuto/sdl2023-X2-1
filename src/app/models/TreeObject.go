package models

type TreeObject struct {
	Elements []*TreeElement
}

type TreeElement struct {
	Meta     string
	Name     string
	ObjectID string
}

func (to *TreeObject) UpdateOrCreateObjectID(name string, objectId string, isDir bool) {
	for _, element := range to.Elements {
		if element.Name == name {
			element.ObjectID = objectId
			return
		}
	}
	meta := ""
	if isDir {
		meta = "040000"
	} else {
		meta = "100644"
	}
	to.Elements = append(to.Elements, &TreeElement{
		Meta:     meta,
		Name:     name,
		ObjectID: objectId,
	})
}
