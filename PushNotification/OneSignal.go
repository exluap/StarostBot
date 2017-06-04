/**
    * @project StarostBot
    * @date 03.06.2017 02:59
    * @author Nikita Zaitsev (exluap) <nickzaytsew@gmail.com>
    * @twitter https://twitter.com/exluap
    *
*/
package PushNotification

import (
	"github.com/exluap/onesignal-sdk-go"
	"StarostBot/tools"
	"fmt"
	"log"
)

var (
	appID string
	appKey string
	userKey string
)
func SendNot(news, url string) {
	client := onesignal.NewClient(nil)
	config := tools.LoadConfig("config.json")

	client.UserKey = config.OneSignal.USERKEY
	client.AppKey = config.OneSignal.APPKEY
	client.RestApiKey = config.OneSignal.RESTAPIKEY
	createNotifications(client,config.OneSignal.APPKEY,news,url)

}

func createNotifications(client *onesignal.Client, appID,news,url string) {
	fmt.Println("### CreateNotifications ###")

	//playerID := "478bef3c-35d6-4a1c-b4f5-71a367784520" // valid
	// playerID := "83823c5f-53ce-4e35-be6a-a3f27e5d838f" // invalid
	if url == "" {
		notificationReq := &onesignal.NotificationRequest{
			AppID:            appID,
			Contents:         map[string]string{"en": news},
			IsIOS:            true,
			IsAndroid:	  true,
			IsAnyWeb:	  true,
			IncludedSegments: []string{"All"},
			Headings:	map[string]string{"en": "Написали какой-то текст от комитета.. Пиздос"},

		}

		createRes, res, err := client.Notifications.Create(notificationReq)
		if err != nil {
			fmt.Printf("--- res:%+v, err:%+v\n", res)
			log.Fatal(err)
		}
		fmt.Printf("--- createRes:%+v\n", createRes)
		fmt.Println()
	} else {
		if tools.Sharing {
			notificationReq := &onesignal.NotificationRequest{
				AppID:            appID,
				Contents:         map[string]string{"en": news},
				IsIOS:            true,
				IsAndroid:	  true,
				IsAnyWeb:	  true,
				IncludedSegments: []string{"All"},
				Headings:         map[string]string{"en": "Тут снова спам с шарингом из группы...."},
				URL:		url,

			}
			createRes, res, err := client.Notifications.Create(notificationReq)
			if err != nil {
				fmt.Printf("--- res:%+v, err:%+v\n", res)
				log.Fatal(err)
			}
			fmt.Printf("--- createRes:%+v\n", createRes)
			fmt.Println()
		}
	}


}
