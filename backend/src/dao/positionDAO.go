package dao

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"model"
	"fmt"
)

const (
	LOCATION_COLLECTION string = "location"
)

var cLocation *mgo.Collection

func InitLocationDAO(database *mgo.Database) {
	fmt.Println("InitLocationDAO")
	cLocation = database.C(LOCATION_COLLECTION)
}

func InsertLocation(userLocation *model.UserLocation) (err error) {
	err = cLocation.Insert(userLocation)
	return
}

func GetAllLocations(filterLocation model.FilterLocation) (locations []model.UserLocation, err error) {
	err = cLocation.Find(filterLocation).All(&locations)
	return
}

