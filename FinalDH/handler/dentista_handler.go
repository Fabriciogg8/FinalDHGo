package handler

import (
	"encoding/json"
	"FinalDH/models"
	"FinalDH/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DentistaHandler struct {
	Service *service.DentistaService
}

func (h *DentistaHandler) GetDentistaByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	dentista, err := h.Service.GetDentistaByID(id)
	if err != nil {
		http.Error(w, "Dentista not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(dentista)
}

func (h *DentistaHandler) CreateDentista(w http.ResponseWriter, r *http.Request) {
	var dentista models.Dentista
	err := json.NewDecoder(r.Body).Decode(&dentista)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = h.Service.CreateDentista(&dentista)
	if err != nil {
		http.Error(w, "Could not create dentista", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dentista)
}

func (h *DentistaHandler) GetAllDentistas(w http.ResponseWriter, r *http.Request) {
	dentistas, err := h.Service.GetAllDentistas()
	if err != nil {
		http.Error(w, "Could not retrieve dentistas", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dentistas)
}

func (h *DentistaHandler) UpdateDentista(w http.ResponseWriter, r *http.Request) {
	var dentista models.Dentista
	err := json.NewDecoder(r.Body).Decode(&dentista)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	dentista.ID = id

	err = h.Service.UpdateDentista(&dentista)
	if err != nil {
		http.Error(w, "Could not update dentista", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dentista)
}

func (h *DentistaHandler) PatchDentista(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var fields map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.Service.PatchDentista(id, fields)
	if err != nil {
		http.Error(w, "Could not update dentista", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *DentistaHandler) DeleteDentista(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := h.Service.DeleteDentista(id)
	if err != nil {
		http.Error(w, "Could not delete dentista", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}