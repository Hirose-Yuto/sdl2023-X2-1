package main

import (
	"main/app"
	"os"
)

func main() {
	gitApp := app.NewApp()
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 2 {
			return
		}
		gitApp.Add(os.Args[2:])
		break
	case "commit":
		if len(os.Args) != 3 {
			return
		}
		gitApp.Commit(os.Args[2])
		break
	}
}
