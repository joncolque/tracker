#!/bin/bash
go version
cd /flowtrace/backend/
export GOPATH=${PWD}
echo
echo GOPATH configurado en $GOPATH
echo
export GOBIN=$GOPATH/bin
echo GOBIN configurado en $GOBIN
cd src/
echo Descargando dependencias, esto puede tardar unos minutos...
go get ./...
echo Dependencias descargadas
echo Iniciando flowtrace-backend
echo
go run main.go