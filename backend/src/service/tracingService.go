package service

import (
	"fmt"
	// "log"
	// "golang.org/x/net/context"
	// firebase "firebase.google.com/go"
	// "firebase.google.com/go/messaging"
	"dao"
	"dto"
	"deliverer"
	"assembler"
)

func Tracing(tracingDevices dto.TracingDevicesDTO) (id_users []string, err error) {
	
	var devices []dto.CompleteDeviceDTO

	//TODO: obtener "app" del JWT_Auth del header y hacer un JWT_Tracker con el nombre de la app y el id_user
	app := "flowtrace"
	JWT_T := "app+id_user" // Este se va a enviar en el payload del mensaje de Firebase. Importar la Lib y generarla

	// Peticion de seguimiento
	for _, id_user := range tracingDevices.Id_users {
		device, err := dao.FindDevice(id_user, app)
		deviceDTO := assembler.ToCompleteDeviceDTO(device)
		if err == nil {
			fmt.Println("token:", deviceDTO.Id_device)
			devices = append(devices, deviceDTO)
		}
	}

	var appPayload dto.AppPayloadDTO
	appPayload.JWT_T = JWT_T
	appPayload.App = app

	deliverer.SetAppPayload(appPayload)
	devicesRegistered, err := deliverer.RegisterDevices(devices, app)

	if err == nil {
		for _, deviceDTO := range devicesRegistered {
			deviceDTO.On_track = true
			device := assembler.FromCompleteDeviceDTO(deviceDTO)
			err := dao.UpdateDevice(&device)
			if err == nil {
				id_users = append(id_users, deviceDTO.Id_user)
			}
		}
	}

	return
}


func StopTracing(tracingDevices dto.TracingDevicesDTO) (id_users []string, err error) {
	
	var devices []dto.CompleteDeviceDTO

	//TODO obtener "app" del JWT_Auth del header y hacer un JWT_Tracker con el nombre de la app y el id_user
	app := "flowtrace"
	// JWT_T := "app+id_user" // Este se va a enviar en el payload del mensaje de Firebase. Importar la Lib y generarla

	for _, id_user := range tracingDevices.Id_users {
		device, err := dao.FindDevice(id_user, app)
		deviceDTO := assembler.ToCompleteDeviceDTO(device)
		if err == nil {
			devices = append(devices, deviceDTO)
		}
	}

	devicesUnRegistered, err := deliverer.UnRegisterDevices(devices, app)

	if err == nil {
		for _, deviceDTO := range devicesUnRegistered {
			deviceDTO.On_track = false
			device := assembler.FromCompleteDeviceDTO(deviceDTO)
			err := dao.UpdateDevice(&device)
			if err == nil {
				id_users = append(id_users, deviceDTO.Id_user)
			}
		}
	}

	return
}