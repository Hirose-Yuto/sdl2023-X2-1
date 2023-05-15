package app

const repoPath = "/exec/.git"

type App struct {
	branchName string
}

func NewApp() *App {
	return &App{branchName: "refs/heads/main"}
}
