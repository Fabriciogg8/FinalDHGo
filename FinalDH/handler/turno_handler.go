package handler

import (
	"encoding/json"
	"FinalDH/model"
	"FinalDH/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TurnoHandler struct {
    Service *service.TurnoService
}

func (h *TurnoHandler) GetAllTurnos(w http.ResponseWriter, r *http.Request) {
    turnos, err := h.Service.GetAllTurnos()
    if err != nil {
        http.Error(w, "Error retrieving turnos", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(turnos)
}

func (h *TurnoHandler) CreateTurno(w http.ResponseWriter, r *http.Request) {
	var turno models.Turno
	if err := json.NewDecoder(r.Body).Decode(&turno); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Service.CreateTurno(&turno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TurnoHandler) GetTurnoByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	turno, err := h.Service.GetTurnoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(turno)
}

func (h *TurnoHandler) UpdateTurno(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var turno models.Turno
	if err := json.NewDecoder(r.Body).Decode(&turno); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	turno.ID = id
	err := h.Service.UpdateTurno(&turno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TurnoHandler) PatchTurno(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var fields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&fields); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Service.PatchTurno(id, fields)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TurnoHandler) DeleteTurno(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := h.Service.DeleteTurno(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TurnoHandler) CreateTurnoByDNIAndMatricula(w http.ResponseWriter, r *http.Request) {
	dni := r.URL.Query().Get("dni")
	matricula := r.URL.Query().Get("matricula")
	var turno models.Turno
	if err := json.NewDecoder(r.Body).Decode(&turno); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Service.CreateTurnoByDNIAndMatricula(dni, matricula, &turno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TurnoHandler) GetTurnoByDNI(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro dni desde los query params
	dni := r.URL.Query().Get("dni")
	if dni == "" {
		http.Error(w, "DNI is required", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para obtener el turno
	turno, err := h.Service.GetTurnoByDNI(dni)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Responder con los datos del turno
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(turno)
}