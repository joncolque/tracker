package assembler

import (
	"dto"
	"model"
)

func ToCompleteDeviceDTO(device model.Device) (completeDeviceDTO dto.CompleteDeviceDTO) {
	completeDeviceDTO.Id_device = device.Id_device
	completeDeviceDTO.Id_user = device.Id_user
	completeDeviceDTO.App = device.App
	completeDeviceDTO.On_track = device.On_track
	return
}

func FromCompleteDeviceDTO(completeDeviceDTO dto.CompleteDeviceDTO) (device model.Device) {
	device.Id_device = completeDeviceDTO.Id_device
	device.Id_user = completeDeviceDTO.Id_user
	device.App = completeDeviceDTO.App 
	device.On_track = completeDeviceDTO.On_track
	return
}

func ToCompleteDevicesDTO(devices []model.Device) (devicesDTO []dto.CompleteDeviceDTO) {
	for _, device := range devices {
		devicesDTO = append(devicesDTO, ToCompleteDeviceDTO(device))
	}
	return
}