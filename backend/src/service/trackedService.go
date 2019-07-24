package service

import (
	"dao"
	"dto"
	"model"

	"assembler"
)

func FindTrackedById(id_tracked string) (trackedDTO dto.TrackedDTO, err error) {
	var tracked model.Tracked
	tracked, err = dao.FindTrackedById(id_tracked)
	trackedDTO = assembler.ToTrackedDTO(tracked)
	return
}

func CreateTracked(trackedDTO *dto.TrackedDTO) (err error) {
	var tracked model.Tracked = assembler.FromTrackedDto(trackedDTO)
	err = dao.CreateTracked(&tracked)
	return
}

func GetAllTracked() (trackedDTOs []dto.TrackedDTO, err error) {
	var tracked []model.Tracked
	tracked, err = dao.GetAllTracked()
	trackedDTOs = assembler.ToTrackedDtos(tracked)
	return
}