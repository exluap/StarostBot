/**
    * @project StarostBot
    * @date 03.06.2017 03:39
    * @author Nikita Zaitsev (exluap) <nickzaytsew@gmail.com>
    * @twitter https://twitter.com/exluap
    *
*/
package tools

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type Configuration struct {
	OneSignal struct {
		APPKEY  string `json:"APP_KEY"`
		USERKEY string `json:"USER_KEY"`
		RESTAPIKEY string `json:"REST_API_KEY"`
	} `json:"OneSignal"`
	VK struct {
		Key string `json:"key"`
		Dialog int `json:"dialog"`
	} `json:"VK"`
}

var (
	Sharing bool = false
)

func LoadConfig(path string) Configuration{

	file,err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Config file missing. ", err)
	}

	var config Configuration

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Fatal("Config parse error: ", err)
	}
	return config
}
