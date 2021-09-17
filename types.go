package main

import "database/sql"

type Modulo struct {
	IdModulo string `json:"IdModulo"`
	Modulo   string `json:"Modulo"`
}

type Tarea struct {
	IdTarea        int            `json:"IdTarea"`
	IdModulo       string         `json:"IdModulo"`
	Unidad         byte           `json:"Unidad"`
	Tipo           string         `json:"Tipo"`
	FechaLimite    string         `json:"FechaLimite"`
	FechaTerminado sql.NullString `json:"FechaTerminado"`
}
