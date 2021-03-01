package collect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"ipPro/config"
	"ipPro/utils"
	"net/http"
	"strings"
)

type domain struct {
	Name string
}

type user struct {
	Domain domain
	Name string
	Passwword string
}

type password struct {
	User user `json:"user"`
}


type domain1 struct {
	Name string `json:"name"`
}

type scope struct {
	Domain domain1 `json:"domain"`
}


type auth struct {
 	Identity identity `json:"identity"`
 	Scope scope `json:"scope"`
 }

 type result struct {
 	Auth auth `json:"auth"`
 }

type identity struct {
	Methods []string `json:"methods"`
	Password password `json:"password"`
}

//func T()  {
//	config.ReadConfig()
//	url := viper.GetString("url")
//	fmt.Println(viper.Get("json"))
//	fmt.Println("url:>", url)
//	//var a auth
//	//a.Identity.Password.User.Name = viper.GetString("data.name")
//	//fmt.Println(a)
//
//	post := result{Auth:auth{Identity:identity{Methods:[]string{"password"},
//		Password:password{
//			User:user{
//				Domain:domain{Name:viper.GetString("data.name"),},
//				Name:viper.GetString("data.name"),
//				Passwword:viper.GetString("data.password")}}},
//	Scope:scope{Domain:domain1{Name:viper.GetString("data.name")}}}}
//
//	fmt.Println(post)
//	//buf,_:= json.MarshalIndent(post, "", "    ") //格式化编码
////	fmt.Println("校验数据",bytes.NewBuffer(buf))
////	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
////	// req.Header.Set("X-Custom-Header", "myvalue")
////	req.Header.Set("Content-Type", "application/json")
////
////	client := &http.Client{}
////	resp, err := client.Do(req)
////	if err != nil{
////		panic(err)
////	}
////	defer resp.Body.Close()
////
////	fmt.Println("status", resp.Status)
////	fmt.Println("response:", resp.Header)
////	body, _ := ioutil.ReadAll(resp.Body)
////	fmt.Println("response Body:", string(body))
////
//}

func GetToken()  {
	config.ReadConfig("cfgfile","config")
	url := viper.GetString("url")
	tokendata := viper.Get("json")

	jsonStr,_ := json.Marshal(tokendata)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json;charset=utf8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "92e77f39-d9c9-4dd0-9eb0-f45356a17fff")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	headData := res.Header
	token := headData["X-Subject-Token"]
	ss := fmt.Sprintf(strings.Join(token, ""))
	fileName := "./test.yaml"
	utils.WriteFile(fileName,"token: "+ss)

	//fmt.Println("v1 type:", reflect.TypeOf(token))
	fmt.Println(ss)
	//headMap := make(map[string]interface{},0)

}


