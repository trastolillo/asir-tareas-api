package main

import "database/sql"

// type Modulo struct {
// 	IdModulo string `json:"id_modulo"`
// 	Modulo   string `json:"modulo"`
// 	Url      string `json:"url"`
// }

// type Unidad struct {
// 	IdModulo string `json:"id_modulo"`
// 	Unidad   byte   `json:"unidad"`
// 	Titulo   string `json:"titulo"`
// 	Url      string `json:"url"`
// }

type Tarea struct {
	IdTarea        int            `json:"id_tarea"`
	IdModulo       string         `json:"id_modulo"`
	Modulo         string         `json:"modulo"`
	Unidad         byte           `json:"unidad"`
	Titulo         string         `json:"titulo"`
	Tipo           string         `json:"tipo"`
	FechaLimite    string         `json:"fecha_limite"`
	FechaTerminado sql.NullString `json:"fecha_terminado"`
}
