package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"ipPro/collect"
	"ipPro/config"
	"net/http"
)

func main() {

	//collect.HostMonitIp()
	//task.Task()

	collect.HostMonitIp()

	//fmt.Println(viper.Get("data.endpoint"))
	//dataSlice1 := make([]string, 0)
	//dataSlice2 := make([]string, 0)
	//dataSlice3 := make([]string, 0)
	////dataSlice4 := make([]string, 0)
	//hostResult := HostMonitIp()
	////huaweiResult := HuaWeiIp()
	//data1 := utils.JSONToMap(hostResult)
	////data2 := utils.JSONToMap(string(huaweiResult))
	////fmt.Println(data2)
	//for _, value := range data1["info"]{
	//	if value.Line == "CM"{
	//		dataSlice1 = append(dataSlice1, value.Ip)
	//		//return dataSlice1
	//	}
	//
	//	if value.Line == "CU"{
	//		dataSlice2 = append(dataSlice2, value.Ip)
	//		//return dataSlice2
	//	}
	//
	//	if value.Line == "CT"{
	//		dataSlice3 = append(dataSlice3,value.Ip)
	//		//return dataSlice3
	//	}
	//
	//}
	//fmt.Println("移动 CM",dataSlice1)
	//fmt.Println("联通 CU",dataSlice2)
	//fmt.Println("电信 CT",dataSlice3)


	//for _, val := range data2["recordsets"]{
	//	//keySlice = append(keySlice, key)
	//	dataSlice4 = append(dataSlice4, val.Line)
	//	fmt.Printf("line=%s",val.Line)
	//	if val.Line == "Yidong"{
	//
	//		//paramsMap := make(map[string]string)
	//		//paramsMap["records"] =HostMoni
	//		//utils.Get(uri,paramsMap,nil)
	//	}
	//}

}





func HuaWeiIp(){

	url := "https://console-intl.huaweicloud.com/dns/rest/v2.1/bundle/feature/line?flat=true&bundle=ultimate"
	config.ReadConfig(" ","test")
	token := viper.GetString("token")
	fmt.Println(token)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Auth-Token",token )
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}


















