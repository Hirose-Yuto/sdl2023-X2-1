package main

import (
	"fmt"
	"main/app/tools"
	"os"
)

func main() {
	oS := tools.NewObjectService("exec/.git")

	readFile(oS)

	//writeFile(oS)

	//sh := sha1.New()
	//io.WriteString(sh, "blob 5\000hello")
	//fmt.Println(hex.EncodeToString(sh.Sum(nil)))
}

func readFile(oS *tools.ObjectService) {
	content, err := oS.ReadBlob("5692df1010ac406517147fa6976ca6e8d15721ec")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)
	fmt.Println([]byte(content))
}

func writeFile(oS *tools.ObjectService) {
	bs, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println(err)
		return
	}
	blob, err := oS.WriteBlob(string(bs))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(blob)
}
