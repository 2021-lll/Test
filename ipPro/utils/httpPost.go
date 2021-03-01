package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type reponseResult struct {
	Total int `json:"total"`
	Info []proxyResult `json:"info"`
}

type proxyResult struct {
	Colo    string    `json:"colo"`
	Ip      string    `json:"ip"`
	Latency string    `json:"latency"`
	Line    string    `json:"line"`
	Loss    string    `json:"loss"`
	Node    string    `json:"node"`
	Speed   string    `json:"speed"`
	Time    string    `json:"time"`
}

func HttpPost(url string, data interface{}, contentType string) string{
	client := &http.Client{Timeout:10*time.Second}
	jsonStr,_ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
		//fmt.Println(err)
	}
	defer resp.Body.Close()

	bodyC,_:= ioutil.ReadAll(resp.Body)


	return string(bodyC)

}

func JSONToMap(str string) map[string][]proxyResult {

	var tempMap = make(map[string][]proxyResult)

	json.Unmarshal([]byte(str), &tempMap)

	return tempMap
}






