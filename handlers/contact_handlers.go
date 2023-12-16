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

func GetAllContact(w http.ResponseWriter, r *http.Request) {
	contacts, err := controllers.GetAllContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(contacts)
}

func CreateNewContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newContact, err := controllers.CreateContact(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newContact)
}

func GetContactByContactID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	contact, err := controllers.GetContactByID(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func UpdateContactByContactID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var updatedContact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&updatedContact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contact, err := controllers.UpdateContactByID(contactID, updatedContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func DeleteContactByContactID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	err = controllers.DeleteContactByID(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
