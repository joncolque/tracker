package service

import (
	"dao"
	"dto"
	"model"
	"fmt"
	"assembler"
)

func FindDevice(id_user string, app string) (completeDevice model.Device, err error) {
	completeDevice, err = dao.FindDevice(id_user, app)
	fmt.Println("completeDevice2",completeDevice)
	return
}

func CreateDevice(deviceToComplete dto.CompleteDeviceDTO) (err error) {
	var device model.Device

	device, err = FindDevice(deviceToComplete.Id_user, deviceToComplete.App)
	device.Id_device = deviceToComplete.Id_device

	if err != nil {
		device.Id_user = deviceToComplete.Id_user
		device.App = deviceToComplete.App
		err = dao.CreateDevice(&device)
	} else {
		err = dao.UpdateDevice(&device)
	}
	return
}

func GetAllDevices() (completeDevicesDTO []dto.CompleteDeviceDTO, err error) {
	var devices []model.Device
	devices, err = dao.GetAllDevices()
	completeDevicesDTO = assembler.ToCompleteDevicesDTO(devices)
	return
}