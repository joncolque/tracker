# Motivación

Este Backend esta desarrollado en Golang e integrada con Firebase Cloud Messaging.
El objetivo de esta aplicación es realizar el seguimiento de dispositivos móviles y persistirlos en una BD en Mongo.

# Documentación

Firebase es un servicio provisto por Google que nos permite comunicarnos con dispositivos móviles. Puede usarse directamente desde su consola (Web) para enviar push notifications a los celulares.

https://console.firebase.google.com/u/0/

Otra alternativa es usarlo en un servidor propio mediante una integración. Esa es la opción que se desarrolla en este proyecto.

La integración está mediante la SDK oficial de Firebase Cloud Messaging para Golang.

https://firebase.google.com/docs/cloud-messaging/server?hl=es-419#firebase-admin-sdk-for-fcm

https://godoc.org/firebase.google.com/go/messaging

# Como agregar Firebase Cloud Messaging a proyecto Android

Ir la consola de Firebase https://console.firebase.google.com/u/0/, settings->general->añadir_aplicación y seguir los pasos.
Nota: Se recomienda generar la clave SHA-1 que te ofrece Firebase como opcional.

# Levantar el proyecto

(Se asume que ya se encuentra configurado todo en la consola de Firebase (Web) y en el cliente Android)
Antes de levantar el proyecto descargar de la consola de Firebase.

setting->service_accounts->generate_new_private_key

eso descarga un archivo .json, debe renombrarse a a *serviceAccountKey.json* y ponerlo en la carpeta backend/
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

- Registrar dispositivo

GET     http://localhost:8080/devices            muestra todos los dispositivos registrados

PUT     http://localhost:8080/devices/:id_user   upsert de id_device (token mobile que se generó en el cliente de firebase)

- Seguimiento

POST    http://localhost:8080/tracing            Inicia o termina el seguimiento sobre los dispositivos especificados

- Ubicaciones

POST    http://localhost:8080/locations          Inserta la ubicación del dispositivo (La request es desde el dispositivo)

GET     http://localhost:8080/locations          Se obtiene todas las ubicaciones guardadas

Para mas detalle ver la colección de Postman *flowtraceGo.postman_collection.json*

# Tokens

Cualquier agente externo que requiera usar este microservicio deberá seguir el siguiente flujo.

- Ubicaciones (POST)

El dispositivo deberá tener en su Header->Authorization un *JWT_T* generado por este microservicio el cual se utiliza para asegurar que el dispositivo es confiable. Este *JWT_T* lo obtendrá de la notificación push silenciosa que se le mandará mediante Firebase.

- Recursos restantes

Cualquier tipo de petición deberá tener en su Header->Authorization un *JWT_Auth* este es generado por el microservicio *Auth* para autentificar la aplicación.

*nota*: Los JWT no se encuentran implementados, pero está el middleware necesario para hacerlo. El microservicio *Auth* aun no se encuentra implementado.

# Configuracion del tiempo de dispatch para obtener las ubicaciones de los dispositivos

El tiempo se configura en los siguiente lugares:

- /backend/src/config/config.dev.yml para DEV.
- /backend/src/config/config.prod.yml para PROD.

Con el siguiente formato:

Unidades "ns", "us" (or "µs"), "ms", "s", "m", "h".
Ejemplo

```yml
dispatch_time: 30s
```

# Links útiles
https://docs.mongodb.com/manual/reference/geojson/