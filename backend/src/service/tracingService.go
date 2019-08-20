package service

import (
	"fmt"
	"model"
	"dao"
	"dto"
	"deliverer"
	"assembler"
)

func Tracing(tracingDevices dto.TracingDevicesDTO) (id_users []string, err error) {
	
	var devices []dto.DeviceDTO

	var deviceFinder model.DeviceFinder
	deviceFinder.App = "flowtrace" //Obtener del JWT_A

	for _, id_user := range tracingDevices.Id_users {
		deviceFinder.Id_user = id_user
		device, err := dao.FindDevice(deviceFinder)
		deviceDTO := assembler.ToDeviceDTO(device)
		if err == nil {
			fmt.Println("token:", deviceDTO.Id_device)

			//TODO: Agregar funcion de JWT que genere el JWT_T a partir de Id_user y deviceDTO
			JWT_T:= deviceDTO.Id_user + deviceDTO.App
			devices = append(devices, deviceDTO)

			err := deliverer.RegisterDevice(deviceDTO, JWT_T)

			if err == nil {
				deviceDTO.On_track = true
				device := assembler.FromDeviceDTO(deviceDTO)
				err := dao.UpdateDevice(deviceFinder,&device)
				if err == nil {
					id_users = append(id_users, deviceDTO.Id_user)
				}
			}
		}
	}

	return
}


func StopTracing(tracingDevices dto.TracingDevicesDTO) (id_users []string, err error) {
	var deviceFinder model.DeviceFinder

	//TODO obtener "app" del JWT_Auth del header
	deviceFinder.App = "flowtrace"

	for _, id_user := range tracingDevices.Id_users {
		deviceFinder.Id_user = id_user
		device, err := dao.FindDevice(deviceFinder)
		deviceDTO := assembler.ToDeviceDTO(device)
		if err == nil {
			deliverer.UnregisterDevice(deviceDTO)
			deviceDTO.On_track = false
			device := assembler.FromDeviceDTO(deviceDTO)
			err := dao.UpdateDevice(deviceFinder, &device)
			if err == nil {
				id_users = append(id_users, deviceDTO.Id_user)
			}
		}
	}

	return
}