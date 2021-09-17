package main

func createTarea(tarea Tarea) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO Agenda (IdModulo, Unidad, Tipo, FechaLimite) VALUES (?, ?, ?, ?) ", tarea.IdModulo, tarea.Unidad, tarea.Tipo, tarea.FechaLimite)
	return err
}

func deleteTarea(idTarea int) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM Agenda WHERE IdTarea=?", idTarea)
	return err
}

func updateTarea(tarea Tarea) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE Agenda SET IdModulo=?, Unidad=?, Tipo=?, FechaLimite=?, FechaTerminado=?", tarea.IdModulo, tarea.Unidad, tarea.Tipo, tarea.FechaLimite, tarea.FechaTerminado)
	return err
}

func getTareas() ([]Tarea, error) {
	tareas := []Tarea{}
	bd, err := getDB()
	if err != nil {
		return tareas, err
	}
	rows, err := bd.Query("SELECT IdTarea, IdModulo, Unidad, Tipo, FechaLimite, FechaTerminado FROM Agenda")
	if err != nil {
		return tareas, err
	}
	for rows.Next() {
		var tarea Tarea
		err = rows.Scan(
			&tarea.IdTarea,
			&tarea.IdModulo,
			&tarea.Unidad,
			&tarea.Tipo,
			&tarea.FechaLimite,
			&tarea.FechaTerminado)
		if err != nil {
			return tareas, err
		}
		tareas = append(tareas, tarea)
	}
	return tareas, nil
}

func getTarea(idTarea int) (Tarea, error) {
	var tarea Tarea
	bd, err := getDB()
	if err != nil {
		return tarea, err
	}
	row := bd.QueryRow("SELECT IdTarea, IdModulo, Unidad, Tipo, FechaLimite, FechaTerminado FROM Agenda WHERE IdTarea=?", idTarea)
	err = row.Scan(
		&tarea.IdTarea,
		&tarea.IdModulo,
		&tarea.Unidad,
		&tarea.Tipo,
		&tarea.FechaLimite,
		&tarea.FechaTerminado)
	if err != nil {
		return tarea, err
	}
	return tarea, nil
}
