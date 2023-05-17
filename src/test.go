package main

import (
	"fmt"
	"main/app/models"
	"main/app/tools"
)

func main() {
	//readBlob()
	//
	//readTree()
	//
	//readCommit()
	//
	//writeFile()
}

func readBlob() {
	oS := tools.NewObjectService("../.git")
	content, err := oS.ReadBlob("5692df1010ac406517147fa6976ca6e8d15721ec")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)
	fmt.Println([]byte(content))
}

func readTree() {
	oS := tools.NewObjectService("../.git")
	content, err := oS.ReadTree("ebf1dff4838af9d9f484831d3c8979738717da56")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, element := range content.Elements {
		fmt.Println(element)
	}
}

func readCommit() {
	oS := tools.NewObjectService("../.git")
	commit, err := oS.ReadCommit("e01e63a00f1e55d84cc6858a08e97ce9e4cfc4ea")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(commit)
}

func readRawObject() {
	oio := tools.NewObjectIO("exec/.git")
	d, err := oio.ReadObject("4b825dc642cb6eb9a060e54bf8d69288fbee4904", models.TREE)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	fmt.Println(string(d))
}

func writeFile() {
	//oS := tools.NewObjectService("../.git")
	//bs, err := os.ReadFile("go.mod")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//blob, err := oS.WriteObject(string(bs))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(blob)
}
