package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"model"
)

const (
	TRACKED_COLLECTION string = "tracked"
)

var cTracked *mgo.Collection

func InitTrackedDAO(database *mgo.Database) {
	cTracked = database.C(TRACKED_COLLECTION)
}

func FindTrackedById(id_tracked string) (tracked model.Tracked, err error) {
	err = cTracked.Find(bson.M{"id_tracked": id_tracked}).One(&tracked)
	return
}

func CreateTracked(tracked *model.Tracked) error {
	return cTracked.Insert(tracked)
}

func GetAllTracked() (tracked []model.Tracked, err error) {
	err = cTracked.Find(bson.M{}).All(&tracked)
	return
}