package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutesForTareas(router *mux.Router) {
	// enableCORS(router)

	router.HandleFunc("/tareas",
		func(w http.ResponseWriter, req *http.Request) {
			tareas, err := getTareas()
			if err == nil {
				respondWithSuccess(tareas, w)
			} else {
				respondWithError(err, w)
			}
		}).Methods(http.MethodGet)

	router.HandleFunc("/tarea/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			idAsString := mux.Vars(r)["id"]
			id, err := stringToInt(idAsString)
			if err != nil {
				respondWithError(err, w)
				return
			}
			tarea, err := getTarea(id)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(tarea, w)
			}
		}).Methods(http.MethodGet)

	router.HandleFunc("/tarea",
		func(w http.ResponseWriter, r *http.Request) {
			var tarea Tarea
			err := json.NewDecoder(r.Body).Decode(&tarea)
			if err != nil {
				respondWithError(err, w)
			} else {
				err := createTarea(tarea)
				if err != nil {
					respondWithError(err, w)
				} else {
					respondWithSuccess(true, w)
				}
			}
		}).Methods(http.MethodPost)

	router.HandleFunc("/tarea",
		func(w http.ResponseWriter, r *http.Request) {
			var tarea Tarea
			err := json.NewDecoder(r.Body).Decode(&tarea)
			if err != nil {
				respondWithError(err, w)
			} else {
				err := updateTarea(tarea)
				if err != nil {
					respondWithError(err, w)
				} else {
					respondWithSuccess(true, w)
				}
			}
		}).Methods(http.MethodPut)

	router.HandleFunc("/tarea/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			idAsString := mux.Vars(r)["id"]
			id, err := stringToInt(idAsString)
			if err != nil {
				respondWithError(err, w)
				return
			}
			err = deleteTarea(id)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}).Methods(http.MethodDelete)
}

// func enableCORS(router *mux.Router) {
// 	router.PathPrefix("/").HandlerFunc(
// 		func(w http.ResponseWriter, req *http.Request) {
// 			w.Header().Set("Acces-Control-Allow-Origin", AllowedCORSDomain)
// 		}).Methods(http.MethodOptions)
// 	router.Use(middlewareCors)
// }

// func middlewareCors(next http.Handler) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, req *http.Request) {
// 			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
// 			w.Header().Set("Access-Control-Allow-Credentials", "true")
// 			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 			next.ServeHTTP(w, req)
// 		})
// }

func respondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}
