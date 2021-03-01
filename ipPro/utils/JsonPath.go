package utils

import (
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
)

func GetBranch(raw []byte)  {
	ipFile, err := jsonpath.Prepare("$..line")

	if err !=nil{
		panic(err)
	}
	var data  interface{}
	if err = json.Unmarshal(raw,&data); err != nil{
		panic(err)
	}
	 out, err := ipFile(data)
	if err != nil{
		panic(err)
	}

	fmt.Println(out)

}



