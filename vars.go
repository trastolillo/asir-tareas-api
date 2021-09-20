package main

import (
	"fmt"
	"os"
)

// var _ = godotenv.Load(".env")

const IsLocal = false

var (
	RemoteConnectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("userBD"),
		os.Getenv("passBD"),
		os.Getenv("hostBD"),
		os.Getenv("nameBD"),
	)
)

const LocalConnectionString = "root:root@tcp(localhost:3306)/Agenda"

// Constantes del esquema de la BBDD
var DatabaseTables = map[string]string{
	"modulos":  "modulos",
	"unidades": "unidades",
	"tareas":   "tareas",
}

var CamposTareas = map[string]string{
	"IdModulo":       "id_modulo",
	"Unidad":         "unidad",
	"Tipo":           "tipo",
	"FechaLimite":    "fecha_limite",
	"FechaTerminado": "fecha_terminado",
}

var CamposModulos = map[string]string{
	"IdModulo": "id_modulo",
	"Modulo":   "modulo",
	"Url":      "url",
}

var CamposUnidades = map[string]string{
	"IdModulo": "id_modulo",
	"Unidad":   "unidad",
	"Titulo":   "titulo",
	"Url":      "url",
}

// const AllowedCORSDomain = "http://localhost"
