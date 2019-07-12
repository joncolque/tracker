package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"model"
)

const (
	SEGUIMIENTO_COLLECTION string = "seguimientos"
)

var cSeguimiento *mgo.Collection

func InitSeguimientoDAO(database *mgo.Database) {
	cSeguimiento = database.C(SEGUIMIENTO_COLLECTION)
}

func FindSeguimientoByName(name string) (seguimiento model.Seguimiento, err error) {
	err = cSeguimiento.Find(bson.M{"name": name}).One(&seguimiento)
	return
}

func FindIdSeguimientoByName(name string) (Id interface{}, err error) {
	var seguimiento model.MongoObject
	err = cSeguimiento.Find(bson.M{"name": name}).One(&seguimiento)
	Id = seguimiento.Id
	return
}

func CreateSeguimiento(seguimiento *model.Seguimiento) error {
	return cSeguimiento.Insert(seguimiento)
}

func GetAllSeguimientos() (seguimientos []model.Seguimiento, err error) {
	err = cSeguimiento.Find(bson.M{}).All(&seguimientos)
	return
}