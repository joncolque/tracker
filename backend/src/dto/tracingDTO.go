package dto

import (
)

type TracingDTO struct {
	Id_tracked		string	`json:"id_tracked" bson:"id_tracked" binding:"required"`
}
