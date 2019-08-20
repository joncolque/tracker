package assembler

import (
	"dto"
	"model"
)

func ToDeviceDTO(device model.Device) (deviceDTO dto.DeviceDTO) {
	deviceDTO.Id_device = device.Id_device
	deviceDTO.Id_user = device.Id_user
	deviceDTO.App = device.App
	deviceDTO.On_track = device.On_track
	return
}

func FromDeviceDTO(deviceDTO dto.DeviceDTO) (device model.Device) {
	device.Id_device = deviceDTO.Id_device
	device.Id_user = deviceDTO.Id_user
	device.App = deviceDTO.App
	device.On_track = deviceDTO.On_track
	return
}

func ToDevicesDTO(devices []model.Device) (devicesDTO []dto.DeviceDTO) {
	for _, device := range devices {
		devicesDTO = append(devicesDTO, ToDeviceDTO(device))
	}
	return
}


func FromUpsertDeviceDTO(device dto.UpsertDeviceDTO) (mDeviceFinder model.DeviceFinder){
	mDeviceFinder.Id_user = device.Id_user
	mDeviceFinder.App = device.App
	return
}