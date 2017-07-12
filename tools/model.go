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
	"os"
	"fmt"
	"strings"
	"net/http"
	"io"
)

type Configuration struct {
	VK struct {
		   Key string `json:"key"`
		   Dialog int64 `json:"dialog"`
		   KeyWord string `json:"keyWord"`
		   MessageWithAttachments string `json:"MessageWithAttachments"`
		   MessageWithoutAttachments string `json:"MessageWithoutAttachments"`
	   } `json:"VK"`
	OneSignal struct {
		   APPKEY string `json:"APP_KEY"`
		   USERKEY string `json:"USER_KEY"`
		   RESTAPIKEY string `json:"REST_API_KEY"`
	   } `json:"OneSignal"`
}

var (
	Sharing bool = false
)

func LoadConfig(path string) Configuration{

	file,err := ioutil.ReadFile(path)

	if err != nil {
		downloadFromUrl("https://raw.githubusercontent.com/exluap/StarostBot/master/config.example.json")
		log.Fatal("Change values in config.example.json and rename file to config.json")
	}

	var config Configuration

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Fatal("Config parse error: ", err)
	}
	return config
}

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]


	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

}
