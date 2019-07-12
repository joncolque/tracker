package assembler

import (
	"dto"
	"model"
)

func FromSeguimientoDto(seguimientoDTO *dto.SeguimientoDTO) (seguimiento model.Seguimiento) {
	seguimiento.Name = seguimientoDTO.Name
	seguimiento.Detail = seguimientoDTO.Detail
	return
}

func ToSeguimientoCompletoDTO(seguimiento model.Seguimiento) (seguimientoDTO dto.SeguimientoCompletoDTO) {
	seguimientoDTO.Name = seguimiento.Name
	seguimientoDTO.Detail = seguimiento.Detail
	return
}

func ToSeguimientoCabeceraDto(seguimiento model.Seguimiento) (seguimientoDTO dto.SeguimientoCabeceraDTO) {
	seguimientoDTO.Name = seguimiento.Name
	seguimientoDTO.Detail = seguimiento.Detail
	return
}

func ToSeguimientoCabeceraDtos(seguimientos []model.Seguimiento) (seguimientosDTO []dto.SeguimientoCabeceraDTO) {
	for _, seguimiento := range seguimientos {
		seguimientosDTO = append(seguimientosDTO, ToSeguimientoCabeceraDto(seguimiento))
	}
	return
}