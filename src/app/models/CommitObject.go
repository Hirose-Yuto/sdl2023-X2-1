package models

type CommitObject struct {
	Tree      string
	Parent    string
	Author    string
	Committer string
	Message   string
}
