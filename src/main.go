package main

import (
	"bufio"
	"fmt"
	"main/app"
	"os"
	"strings"
)

func main() {
	gitApp, err := app.NewApp()
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

		switch strings.Split(arg, " ")[0] {
		case "add":
			if err := gitApp.Add(strings.Split(arg, " ")[1:]); err != nil {
				fmt.Println(err)
			}
			break
		case "commit":
			if err := gitApp.Commit(arg[7:]); err != nil {
				fmt.Println(err)
			}
			break
		}
	}
}
