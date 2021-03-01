package collect

import (
	"fmt"
	"io/ioutil"
	"ipPro/utils"
	"strings"
)

func HostMonitIp() string {
	//keySlice := make([]int, 0)
	data := make(map[string]interface{})

	data["key"] = "iDetkOys"
	uri := "https://api.hostmonit.com/get_optimization_ip"
	result := utils.HttpPost(uri,data,"application/json")

	//var jsonMap map[string]interface{}
	//jsonMap = make(map[string]interface{})
	//err := json.Unmarshal([]byte(result), &jsonMap)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	//return string(bodyC)

	//
	dataSlice1 := make([]string, 0)
	dataSlice2 := make([]string, 2)
	dataSlice3 := make([]string, 3)
	//dataSlice2 := make([]string, 0)
	//dataSlice3 := make([]string, 0)
	//return string(result)
	dataresult := utils.JSONToMap(result)
	//
	for _, value := range dataresult["info"]{
		if value.Line == "CM"{

			dataSlice1 = append(dataSlice1,value.Ip)
			//return dataSlice1
		}

		if value.Line == "CU"{
			dataSlice2 = append(dataSlice2, value.Ip)
			//return dataSlice2
		}

		if value.Line == "CT"{
			dataSlice3 = append(dataSlice3,value.Ip)
			//return dataSlice3
		}

	}



	//stringByte := "\x00" + strings.Join(dataSlice1, "\x20\x00") // x20 = space and x00 = null
	stringByte1 := strings.Join(dataSlice1, "  ") // x20 = space and x00 = null
	stringByte2 := strings.Join(dataSlice2, "  ")
	stringByte3 := strings.Join(dataSlice3, "  ")


	filepath := "./ip"
	//ipMap := make(map[string]string, 0)

	utils.WriteFile("ip","CM:  "+stringByte1+"CU:"+stringByte2+"CT:"+stringByte3)


	data1, err := ioutil.ReadFile(filepath)
	if err !=nil {
		panic(err)
	}
	fmt.Println(data1)
	fmt.Println(string(data1))



	//fmt.Println("移动 CM",dataSlice1)
	//fmt.Println("联通 CU",dataSlice2)
	//fmt.Println("电信 CT",dataSlice3)
	return string(data1)
}
