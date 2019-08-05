package dto

import (
)

type AppPayloadDTO struct {
	JWT_T	string	`json:"jwt_t" bson:"jwt_t"`
	App	string	`json:"app" bson:"app"`
}