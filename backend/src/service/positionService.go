package service

import (
	"dao"
	"dto"
	"assembler"
	"model"
)

func InsertLocation(location dto.UserLocationDTO) (err error) {
	var mLocation model.UserLocation
	mLocation = assembler.FromUserLocationDTO(location)
	err = dao.InsertLocation(&mLocation)
	return
}

func GetAllLocations(filterLocation dto.FilterLocationDTO) (usersLocationsDTO []dto.UserLocationDTO, err error) {
	var mFilterLocation model.FilterLocation
	var mUserLocations []model.UserLocation
	
	mFilterLocation = assembler.FromFilterLocationDTO(filterLocation)
	mUserLocations, err = dao.GetAllLocations(mFilterLocation)
	usersLocationsDTO = assembler.ToUsersLocationsDTO(mUserLocations)
	return
}