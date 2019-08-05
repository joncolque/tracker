package deliverer

import (
	"dto"
	"fmt"
	"log"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"time"
)

var (
	app 	*firebase.App
	ctx 	context.Context
	message *messaging.Message
)

func InitProcess(){
	fmt.Println("Deliverer InitProcess")
	message = &messaging.Message{
		Data: map[string]string{
				"dummy-data": "dummy-data ",
		},
		Topic: "dummy-application",
		Notification: &messaging.Notification{
			Title: "Hola!",
			Body:  "Como estas?",
		},
	}
	err := ConnectService()
	if err != nil {
		fmt.Println("there was a problem connecting to the firebase service")
	} else {
		fmt.Println("Connection with Firebase was succesfully")
		go Sender()
	}
}

func Sender(){
	for {
		time.Sleep(30 * time.Second)
		if message != nil{
			client := GetClient()
			response, err := client.Send(ctx, message)
			fmt.Println("Send Response", response)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Println("No hay mensajes para despachar")
		}
  	}
}

func GetClient() (client *messaging.Client){
	ctx = context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	return
}

func ConnectService() (err error){
	fmt.Println("connecting with firebase")
	app, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	fmt.Println("connect to firebase")
	return
}

func RegisterDevices(devices []dto.CompleteDeviceDTO, topic string) (registeredDevices []dto.CompleteDeviceDTO, err error){
	var tokenDevices []string
	for _, device := range devices {
		tokenDevices = append(tokenDevices, device.Id_device)
		registeredDevices = append(registeredDevices, device)
	}
	client := GetClient()
	response, err := client.SubscribeToTopic(ctx, tokenDevices, topic)
	return
}

func UnRegisterDevices(devices []dto.CompleteDeviceDTO, topic string) (registeredDevices []dto.CompleteDeviceDTO, err error){
	var tokenDevices []string
	for _, device := range devices {
		tokenDevices = append(tokenDevices, device.Id_device)
		registeredDevices = append(registeredDevices, device)
	}
	client := GetClient()
	response, err := client.UnsubscribeFromTopic(ctx, tokenDevices, topic)
	//TODO: Ver si se pude tomar de la response los tokens subscritos.
	return
}

func SetAppPayload(appPayload dto.AppPayloadDTO){
	fmt.Println("Set app payload")
	message = &messaging.Message{
			Data: map[string]string{
					"JWT_T": appPayload.JWT_T,
			},
			Topic: appPayload.App,
			Notification: &messaging.Notification{
				Title: "Hola!",
				Body:  "Como estas?",
			},
	}
}
