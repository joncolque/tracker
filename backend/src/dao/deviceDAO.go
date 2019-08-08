package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"model"
)

const (
	DEVICE_COLLECTION string = "device"
)

var cDevice *mgo.Collection

func InitDeviceDAO(database *mgo.Database) {
	cDevice = database.C(DEVICE_COLLECTION)
}

func FindDevice(id_user string, app string) (device model.Device, err error) {
	err = cDevice.Find(
			bson.M{
				"id_user": id_user, 
				"app": app,
			},
		).One(&device)
	return
}

func CreateDevice(device *model.Device) error {
	return cDevice.Insert(device)
}

func UpdateDevice(device *model.Device) error {
	return cDevice.Update(
		bson.M{
			"id_user":device.Id_user,
			"app":device.App,
		},
		device)
}

func GetAllDevices() (devices []model.Device, err error) {
	err = cDevice.Find(bson.M{}).All(&devices)
	return
}