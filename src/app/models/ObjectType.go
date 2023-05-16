package models

type ObjectType int64

const (
	BLOB ObjectType = iota
	COMMIT
	TREE
)

func ToString(objectType ObjectType) string {
	switch objectType {
	case BLOB:
		return "blob"
	case COMMIT:
		return "commit"
	case TREE:
		return "tree"
	}
	return ""
}
