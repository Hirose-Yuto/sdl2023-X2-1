package models

type ObjectType int64

const (
	BLOB ObjectType = iota
	COMMIT
	TREE
)

func (o ObjectType) ToString() string {
	switch o {
	case BLOB:
		return "blob"
	case COMMIT:
		return "commit"
	case TREE:
		return "tree"
	}
	return ""
}
