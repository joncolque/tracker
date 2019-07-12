package service

import (
	"dao"
	"dto"
	"model"

	"assembler"
)

func FindSeguimientoByName(name string) (seguimientoCompletoDTO dto.SeguimientoCompletoDTO, err error) {
	var seguimiento model.Seguimiento
	seguimiento, err = dao.FindSeguimientoByName(name)
	seguimientoCompletoDTO = assembler.ToSeguimientoCompletoDTO(seguimiento)
	return
}

func CreateSeguimiento(seguimientoDTO *dto.SeguimientoDTO) (err error) {
	var seguimiento model.Seguimiento = assembler.FromSeguimientoDto(seguimientoDTO)
	err = dao.CreateSeguimiento(&seguimiento)
	return
}

func GetAllSeguimientos() (seguimientosDTO []dto.SeguimientoCabeceraDTO, err error) {
	var seguimientos []model.Seguimiento
	seguimientos, err = dao.GetAllSeguimientos()
	seguimientosDTO = assembler.ToSeguimientoCabeceraDtos(seguimientos)
	return
}