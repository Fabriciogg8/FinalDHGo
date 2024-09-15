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

    // Middleware de autenticación
    r.Use(middleware.AuthMiddleware)

    http.ListenAndServe(":8080", r)
}