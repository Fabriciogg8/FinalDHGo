package main

import (
    "FinalDH/config"
    "FinalDH/handler"
    "FinalDH/middleware"
    "FinalDH/repository"
    "FinalDH/service"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    config.ConnectDB()

    dentistaRepo := &repository.DentistaRepository{DB: config.DB}
    dentistaService := &service.DentistaService{Repo: dentistaRepo}
    dentistaHandler := &handler.DentistaHandler{Service: dentistaService}

    pacienteRepo := &repository.PacienteRepository{DB: config.DB}
    pacienteService := &service.PacienteService{Repo: pacienteRepo}
    pacienteHandler := &handler.PacienteHandler{Service: pacienteService}

    turnoRepo := &repository.TurnoRepository{DB: config.DB}
	turnoService := &service.TurnoService{Repo: turnoRepo}
	turnoHandler := &handler.TurnoHandler{Service: turnoService}

    r := mux.NewRouter()

    // Rutas de Dentista
    r.HandleFunc("/dentistas", dentistaHandler.GetAllDentistas).Methods("GET")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.GetDentistaByID).Methods("GET")
    r.HandleFunc("/dentistas", dentistaHandler.CreateDentista).Methods("POST")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.UpdateDentista).Methods("PUT")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.PatchDentista).Methods("PATCH")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.DeleteDentista).Methods("DELETE")

    // Rutas de Paciente
    r.HandleFunc("/pacientes", pacienteHandler.GetAllPacientes).Methods("GET")
    r.HandleFunc("/pacientes/{id}", pacienteHandler.GetPacienteByID).Methods("GET")
    r.HandleFunc("/pacientes", pacienteHandler.CreatePaciente).Methods("POST")
    r.HandleFunc("/pacientes/{id}", pacienteHandler.UpdatePaciente).Methods("PUT")
    r.HandleFunc("/pacientes/{id}", pacienteHandler.PatchPaciente).Methods("PATCH")
    r.HandleFunc("/pacientes/{id}", pacienteHandler.DeletePaciente).Methods("DELETE")

    // Rutas de Turno
    r.HandleFunc("/turnos", turnoHandler.GetAllTurnos).Methods("GET")
	r.HandleFunc("/turnos", turnoHandler.CreateTurno).Methods("POST")
	r.HandleFunc("/turnos/{id}", turnoHandler.GetTurnoByID).Methods("GET")
	r.HandleFunc("/turnos/{id}", turnoHandler.UpdateTurno).Methods("PUT")
	r.HandleFunc("/turnos/{id}", turnoHandler.PatchTurno).Methods("PATCH")
	r.HandleFunc("/turnos/{id}", turnoHandler.DeleteTurno).Methods("DELETE")
	r.HandleFunc("/turnos/by-dni-matricula", turnoHandler.CreateTurnoByDNIAndMatricula).Methods("POST")
	r.HandleFunc("/turnos/by-dni", turnoHandler.GetTurnoByDNI).Methods("GET")

    // Middleware de autenticación
    r.Use(middleware.AuthMiddleware)

    http.ListenAndServe(":8080", r)
}