package assembler

import (
	"dto"
	"model"
)

func FromTrackedDto(trackedDTO *dto.TrackedDTO) (tracked model.Tracked) {
	tracked.Id_tracked = trackedDTO.Id_tracked
	tracked.Token_tracked = trackedDTO.Token_tracked
	return
}

func ToTrackedDTO(tracked model.Tracked) (trackedDTO dto.TrackedDTO) {
	trackedDTO.Id_tracked = tracked.Id_tracked
	trackedDTO.Token_tracked = tracked.Token_tracked
	return
}

func ToTrackedDtos(trackedDevices []model.Tracked) (trackedDTOs []dto.TrackedDTO) {
	for _, tracked := range trackedDevices {
		trackedDTOs = append(trackedDTOs, ToTrackedDTO(tracked))
	}
	return
}