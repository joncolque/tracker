package model

import (
)

type Device struct {
	Id_device		string	`json:"id_device" bson:"id_device"`
	Id_user			string	`json:"id_user" bson:"id_user"`
	App				string	`json:"app" bson:"app"`
	On_track		bool    `json:"on_track" bson:"on_track"`
}