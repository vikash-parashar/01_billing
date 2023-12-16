package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vikash-parashar/01_billing/controllers"
	"github.com/vikash-parashar/01_billing/models"
)

func GetAllInsuredHandler(w http.ResponseWriter, r *http.Request) {
	insuredList, err := controllers.GetAllInsured()
	if err != nil {
		http.Error(w, "Failed to fetch insured list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insuredList)
}

func CreateInsuredHandler(w http.ResponseWriter, r *http.Request) {
	var insured models.Insured
	if err := json.NewDecoder(r.Body).Decode(&insured); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, err := controllers.CreateInsured(insured)
	if err != nil {
		http.Error(w, "Failed to create insured", http.StatusInternalServerError)
		return
	}

	insured.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insured)
}

func GetInsuredByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	insured, err := controllers.GetInsuredByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Insured not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get insured", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insured)
}

func UpdateInsuredByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	var insured models.Insured
	if err := json.NewDecoder(r.Body).Decode(&insured); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = controllers.UpdateInsuredByID(id, insured)
	if err != nil {
		http.Error(w, "Failed to update insured", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteInsuredByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	err = controllers.DeleteInsuredByID(id)
	if err != nil {
		http.Error(w, "Failed to delete insured", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
