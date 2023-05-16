package models

type TreeObject struct {
	Elements []*TreeElement
}

type TreeElement struct {
	Meta     string
	Name     string
	ObjectID string
}
