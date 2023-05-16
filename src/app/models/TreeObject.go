package models

type TreeObject struct {
	Elements []*TreeElement
}

type TreeElement struct {
	Meta     string
	Name     string
	ObjectID string
}

func (to *TreeObject) UpdateObjectID(name string, objectId string) {
	for _, element := range to.Elements {
		if element.Name == name {
			element.ObjectID = objectId
		}
	}
}
