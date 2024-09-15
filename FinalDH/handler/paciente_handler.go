package handler

import (
	"encoding/json"
	"FinalDH/model"
	"FinalDH/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PacienteHandler struct {
	Service *service.PacienteService
}

func (h *PacienteHandler) GetPacienteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	paciente, err := h.Service.GetPacienteByID(id)
	if err != nil {
		http.Error(w, "Paciente not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(paciente)
}

func (h *PacienteHandler) CreatePaciente(w http.ResponseWriter, r *http.Request) {
	var paciente models.Paciente
	err := json.NewDecoder(r.Body).Decode(&paciente)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = h.Service.CreatePaciente(&paciente)
	if err != nil {
		http.Error(w, "Could not create paciente", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(paciente)
}

func (h *PacienteHandler) GetAllPacientes(w http.ResponseWriter, r *http.Request) {
	pacientes, err := h.Service.GetAllPacientes()
	if err != nil {
		http.Error(w, "Could not retrieve pacientes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pacientes)
}

func (h *PacienteHandler) UpdatePaciente(w http.ResponseWriter, r *http.Request) {
	var paciente models.Paciente
	err := json.NewDecoder(r.Body).Decode(&paciente)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	paciente.ID = id

	err = h.Service.UpdatePaciente(&paciente)
	if err != nil {
		http.Error(w, "Could not update paciente", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(paciente)
}

func (h *PacienteHandler) PatchPaciente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var fields map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.Service.PatchPaciente(id, fields)
	if err != nil {
		http.Error(w, "Could not update paciente", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *PacienteHandler) DeletePaciente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := h.Service.DeletePaciente(id)
	if err != nil {
		http.Error(w, "Could not delete paciente", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}