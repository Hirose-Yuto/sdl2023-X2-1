package main

import (
	"fmt"
	"main/app/tools"
)

func main() {
	oS := tools.NewObjectService("../.git")

	//readBlob(oS)

	//readTree(oS)

	readCommit(oS)

	//writeFile(oS)

}

func readBlob(oS *tools.ObjectService) {
	content, err := oS.ReadBlob("5692df1010ac406517147fa6976ca6e8d15721ec")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)
	fmt.Println([]byte(content))
}

func readTree(oS *tools.ObjectService) {
	content, err := oS.ReadTree("7acdc43c587e2e299750279ac5191dc4498b0e26")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, element := range content.Elements {
		fmt.Println(element)
	}
}

func readCommit(oS *tools.ObjectService) {
	commit, err := oS.ReadCommit("e01e63a00f1e55d84cc6858a08e97ce9e4cfc4ea")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(commit)
}

func writeFile(oS *tools.ObjectService) {
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
