# Motivación

Este Backend esta desarrollado en Golang e integrada con Firebase Cloud Messaging.
El objetivo de esta aplicación es realizar el seguimiento de dispositivos móviles y persistirlos en una BD en Mongo.

# Documentación

Firebase es un servicio provisto por Google que nos permite comunicarnos con dispositivos móviles. Puede usarse directamente desde su consola para enviar push notifitations a los celulares.

https://console.firebase.google.com/u/0/

O tambien puede usarse en un servidor propio mediante una integración.
La integración está mediante la SDK oficial provista por Google para Golang.

https://firebase.google.com/docs/cloud-messaging/server?hl=es-419#firebase-admin-sdk-for-fcm

https://godoc.org/firebase.google.com/go/messaging

## Como agregar Firebase Cloud Messaging a proyecto Android

Ir la consola de Firebase https://console.firebase.google.com/u/0/, settings->general->añadir_aplicación y seguir los pasos.
Nota: Se recomienda generar la clave SHA-1 que te ofrece Firebase como opcional.

# Levantar el proyecto

(Se asume que ya se encuentra configurado todo el servidor de firebase desde su consola)
Antes de levantar el proyecto descargar de la consola de Firebase
setting->service_accounts->generate_new_private_key
eso descarga archivo .json y renombrarlo a *serviceAccountKey.json* y ponerlo en la carpeta backend/
(Nunca subir a git este archivo!)

Para levantar el proyecto la primera vez se debe hacer es lo siguiente:

```bash
docker-compose up -d
docker exec -it flowtrace-backend bash
cd flowtrace/backend
export GOPATH=${PWD}
export GOBIN=$GOPATH/bin
export GOOGLE_APPLICATION_CREDENTIALS=/flowtrace/backend/serviceAccountKey.json
cd src
go get ./...
go run main.go  #Luego solo con este comando levanta.
```

# Endpoints

Servirá en http://localhost:8080

TODO:

- Registrar dispositivo
- Iniciar seguimiento
- Guardar posición

# Tokens

TODO: flujo de tokens de firebase, e internos de CyS.