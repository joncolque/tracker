package dto

import (
)

type SignatureLocationDTO struct {
	Longitude	float64			`json:"longitude" bson:"longitude" binding:"required"`
	Latitude	float64		`json:"latitude" bson:"latitude" binding:"required"`
	Timestamp	string		`json:"timestamp" bson:"timestamp" binding:"required"`
}

type SignatureFilterLocationDTO struct {
	//TODO: ver como bindear dinamicamente cuanto se agreguen mas query params
	Id_user		string		`form:"id_user"`
}

type UserLocationDTO struct {
	Id_user		string		`json:"id_user"`
	App 		string 		`json:"app"`
	Location 	LocationDTO `json:"Location"`
}

type LocationDTO struct {
	Longitude	float64		`json:"longitude"	bson:"longitude"`
	Latitude	float64		`json:"latitude"	bson:"latitude"`
	Timestamp	string		`json:"timestamp"	bson:"timestamp"`
}

type FilterLocationDTO struct {
	//TODO: ver como bindear dinamicamente cuanto se agreguen mas query params
	Id_user		string		`form:"id_user"`
}