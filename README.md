# statistics

Se requiere del servicio: /credit/search_user

Servicio: POST: /statistics
	- Requiere de un token (Autoenthication Beaer token)
	- Antes de ejecutar el servicio valida que tenga cceso.
	- Archivo de configuracion configDB.json, configHttp.json y configMsSarchUsr.json

Responde en caso de exito 200 y un json:
 {
	"asignaciones exitosas": 57,
	"asignaciones no exitosas": 20,
	"asignaciones realizadas": 77,
	"promedio asignaciones exitosas": 7.125,
	"promedio asignaciones no exitosas": 6.6666665
}
 
Responde en caso de error:
 {
	"code": "statistics_microservice",
	"message": "MEnsaje a detalle",
	"status_code": 400
 }

