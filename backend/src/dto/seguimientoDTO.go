package dto

import (
)

type SeguimientoDTO struct {
	Name      string           `json:"name" bson:"name" binding:"required"`
	Detail    string           `json:"detail" bson:"detail"`
}

type SeguimientoCompletoDTO struct {
	Name      string           `json:"name"`
	Detail    string           `json:"detail"`
}

type SeguimientoCabeceraDTO struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}
