package dto

import (
)

type TracingDevicesDTO struct {
	Id_users		[]string	`json:"id_users" bson:"id_users" binding:"required"`
	Follow			bool		`json:"follow" bson:"follow" binding:"exists"`
}
