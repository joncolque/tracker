package service

import (
	"fmt"
	"log"
	"github.com/maddevsio/fcm"
	"github.com/spf13/viper"
)

func InitTracing(id_tracked string) {
	serverKey := viper.GetString("apiToken")
	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}
	c := fcm.NewFCM(serverKey)
	token := id_tracked
	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  []string{token},
		ContentAvailable: false,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: "Hello2",
			Body:  "World",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)
}
