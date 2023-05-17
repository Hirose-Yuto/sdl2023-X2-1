package main

import (
	"bufio"
	"fmt"
	"main/app/commands"
	"os"
	"strings"
)

func main() {
	gitApp, err := commands.NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("% ")
		if !scanner.Scan() {
			fmt.Println("\nexit...")
			return
		}
		arg := scanner.Text()
		if strings.Split(arg, " ")[0] != "git" {
			fmt.Printf("%s is not recognized as a command\n", strings.Split(arg, " ")[0])
		}

		switch strings.Split(arg, " ")[1] {
		case "init":
			if err := gitApp.Init(); err != nil {
				fmt.Println(err)
				return
			}
			break
		case "add":
			if err := gitApp.Add(strings.Split(arg, " ")[2:]); err != nil {
				fmt.Println(err)
			}
			break
		case "commit":
			if err := gitApp.Commit(strings.Trim(arg[11:], "\"\"")); err != nil {
				fmt.Println(err)
			}
			break
		case "status":
			if err := gitApp.Status(); err != nil {
				fmt.Println(err)
				return
			}
			break
		case "ls-files":
			if err := gitApp.LsFiles(); err != nil {
				fmt.Println(err)
				return
			}
			break
		case "log":
			if err := gitApp.Log(); err != nil {
				fmt.Println(err)
				return
			}
			break
		}
	}
}
