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
	"github.com/exluap/StarostBot/tools"
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
	config := tools.LoadConfig("config.example.json")

	client.UserKey = config.OneSignal.USERKEY
	client.AppKey = config.OneSignal.APPKEY
	client.RestApiKey = config.OneSignal.RESTAPIKEY
	createNotifications(client,config.OneSignal.APPKEY,news,url, config.VK.MessageWithAttachments,config.VK.MessageWithoutAttachments)

}

func createNotifications(client *onesignal.Client, appID,news,url string, attachments string, withoutattachments string) {
	fmt.Println("### CreateNotifications ###")
	if url == "" {
		notificationReq := &onesignal.NotificationRequest{
			AppID:            appID,
			Contents:         map[string]string{"en": news},
			IsIOS:            true,
			IsAndroid:	  true,
			IsAnyWeb:	  true,
			IncludedSegments: []string{"All"},
			Headings:	map[string]string{"en": withoutattachments},

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
				Headings:         map[string]string{"en": attachments},
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
