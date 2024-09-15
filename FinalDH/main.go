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

    r := mux.NewRouter()

    // Rutas de Dentista
    r.HandleFunc("/dentistas", dentistaHandler.GetAllDentistas).Methods("GET")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.GetDentistaByID).Methods("GET")
    r.HandleFunc("/dentistas", dentistaHandler.CreateDentista).Methods("POST")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.UpdateDentista).Methods("PUT")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.PatchDentista).Methods("PATCH")
    r.HandleFunc("/dentistas/{id}", dentistaHandler.DeleteDentista).Methods("DELETE")

    // Middleware de autenticación
    r.Use(middleware.AuthMiddleware)

    http.ListenAndServe(":8080", r)
}