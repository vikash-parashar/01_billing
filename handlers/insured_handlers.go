// handlers.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vikash-parashar/01_billing/controllers"
	"github.com/vikash-parashar/01_billing/models"
)

func GetAllInsured(w http.ResponseWriter, r *http.Request) {
	insureds, err := controllers.GetAllInsureds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(insureds)
}

func CreateNewInsured(w http.ResponseWriter, r *http.Request) {
	var insured models.Insured
	if err := json.NewDecoder(r.Body).Decode(&insured); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newInsured, err := controllers.CreateInsured(insured)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newInsured)
}

func GetInsuredByInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	insured, err := controllers.GetInsuredByID(insuredID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(insured)
}

func UpdateInsuredByInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	var updatedInsured models.Insured
	if err := json.NewDecoder(r.Body).Decode(&updatedInsured); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insured, err := controllers.UpdateInsuredByID(insuredID, updatedInsured)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(insured)
}

func DeleteInsuredByInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	err = controllers.DeleteInsuredByID(insuredID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllInsuredByContactID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	insureds, err := controllers.GetAllInsuredsByContactID(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(insureds)
}

func CreateNewInsuredByContactID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var insured models.Insured
	if err := json.NewDecoder(r.Body).Decode(&insured); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newInsured, err := controllers.CreateInsuredByContactID(contactID, insured)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newInsured)
}

func GetInsuredByContactAndInsuredID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	insured, err := controllers.GetInsuredByContactAndInsuredID(contactID, insuredID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(insured)
}

func UpdateInsuredByContactAndInsuredID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	var updatedInsured models.Insured
	if err := json.NewDecoder(r.Body).Decode(&updatedInsured); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insured, err := controllers.UpdateInsuredByContactAndInsuredID(contactID, insuredID, updatedInsured)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(insured)
}

func DeleteInsuredByContactAndInsuredID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	err = controllers.DeleteInsuredByContactAndInsuredID(contactID, insuredID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
