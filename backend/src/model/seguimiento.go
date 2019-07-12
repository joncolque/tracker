package model

import (
)

type Seguimiento struct {
	Name      string      `json:"name" bson:"name" binding:"required"`
	Detail    string      `json:"detail" bson:"detail"`
}

type SeguimientoName struct{
	Name      string      `json:"name" bson:"name" binding:"required"`	
}