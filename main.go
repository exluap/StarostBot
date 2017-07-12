/**
    * @project StarostBot
    * @date 03.06.2017 02:14
    * @author Nikita Zaitsev (exluap) <nickzaytsew@gmail.com>
    * @twitter https://twitter.com/exluap
    *
*/
package main

import (
	"github.com/urShadow/go-vk-api"
	"log"
	"fmt"
	"strings"
	"github.com/exluap/StarostBot/tools"
	"github.com/exluap/StarostBot/PushNotification"
	"os"
)

const VK_URL = "https://vk.com/wall"

func main() {

	if os.Args[1] == "true" {
		tools.Sharing = true
	}


	config:=tools.LoadConfig("config.example.json")

	api := vk.New("ru")

	err := api.Init(config.VK.Key)

	if err != nil {
		log.Fatal("Error while init: ", err)
	}


	api.OnNewMessage(func(msg *vk.LPMessage) {
		if msg.Flags&vk.FlagMessageOutBox ==0 {
			if msg.FromID == config.VK.Dialog && strings.HasPrefix(msg.Text,config.VK.KeyWord)  {

					if len(msg.Attachments)>1 {
						res := strings.Split(msg.Text,config.VK.KeyWord)

						if res[1] == "" {
							PushNotification.SendNot(res[1],VK_URL + msg.Attachments["attach1"])
						} else {
							PushNotification.SendNot(res[1], VK_URL + msg.Attachments["attach1"])
						}
					} else {
						res := strings.Split(msg.Text,config.VK.KeyWord)
						fmt.Println(res[1])

						PushNotification.SendNot(res[1], "")

					}

			}
		}
	})

	api.RunLongPoll()
}



