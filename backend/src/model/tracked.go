package model

import (
)

type Tracked struct {
	Id_tracked		string	`json:"id_tracked" bson:"id_tracked" binding:"required"`
	Token_tracked	string	`json:"token_tracked" bson:"token_tracked"`
}