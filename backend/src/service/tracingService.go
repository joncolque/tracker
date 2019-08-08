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
	app:="flowtrace"	//Obtener del JWT_A

	// Peticion de seguimiento
	for _, id_user := range tracingDevices.Id_users {
		device, err := dao.FindDevice(id_user, app)
		deviceDTO := assembler.ToCompleteDeviceDTO(device)
		if err == nil {
			fmt.Println("token:", deviceDTO.Id_device)
			//TODO: Agregar funcion de JWT que lo genere
			JWT_T:= deviceDTO.Id_user + deviceDTO.App
			devices = append(devices, deviceDTO)

			err := deliverer.RegisterDevice(deviceDTO, JWT_T)

			if err == nil {
				deviceDTO.On_track = true
				device := assembler.FromCompleteDeviceDTO(deviceDTO)
				err := dao.UpdateDevice(&device)
				if err == nil {
					id_users = append(id_users, deviceDTO.Id_user)
				}
			}
		}
	}

	return
}


func StopTracing(tracingDevices dto.TracingDevicesDTO) (id_users []string, err error) {
	
	//TODO obtener "app" del JWT_Auth del header y hacer un JWT_Tracker con el nombre de la app y el id_user
	app := "flowtrace"

	for _, id_user := range tracingDevices.Id_users {
		device, err := dao.FindDevice(id_user, app)
		deviceDTO := assembler.ToCompleteDeviceDTO(device)
		if err == nil {
			deliverer.UnregisterDevice(deviceDTO)
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