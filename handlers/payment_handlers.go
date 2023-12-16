package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vikash-parashar/01_billing/controllers"
	"github.com/vikash-parashar/01_billing/models"
)

// CreatePaymentMethodHandler handles the creation of a new payment method for a given contact ID.
func CreatePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
	contactID := chi.URLParam(r, "id")

	// Parse the request body to get the payment method details
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate and handle the payment method creation
	res, err := controllers.CreatePaymentMethod(contactID, payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	log.Printf("New Payment Inserted With Id : %d\n", id)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Payment method created successfully"))
}

func GetAllPaymentMethodsForContactHandler(w http.ResponseWriter, r *http.Request) {
	contactID := chi.URLParam(r, "id")

	// Retrieve payment methods from the database
	paymentMethods, err := controllers.GetAllPaymentMethodsForContact(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the payment methods to JSON
	response, err := json.Marshal(paymentMethods)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetPaymentMethodHandlerByMethodID retrieves a specific payment method for a given contact and payment method ID.
func GetPaymentMethodHandlerByMethodID(w http.ResponseWriter, r *http.Request) {
	contactID := chi.URLParam(r, "id")
	methodID := chi.URLParam(r, "methodId")

	// Retrieve the payment method from the database
	paymentMethod, err := controllers.GetPaymentMethodByMethodID(contactID, methodID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the payment method to JSON
	response, err := json.Marshal(paymentMethod)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// DeletePaymentMethodHandler deletes a specific payment method for a given contact and payment method ID.
func DeletePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
	contactID := chi.URLParam(r, "id")
	methodID := chi.URLParam(r, "methodId")

	// Delete the payment method from the database
	err := controllers.DeletePaymentMethod(contactID, methodID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdatePaymentMethodHandler updates a specific payment method for a given contact and payment method ID.
func UpdatePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
	contactID := chi.URLParam(r, "id")
	methodID := chi.URLParam(r, "methodId")

	// Parse the request body to get the updated payment method details
	var updatedPayment models.Payment
	err := json.NewDecoder(r.Body).Decode(&updatedPayment)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the payment method in the database
	err = controllers.UpdatePaymentMethod(contactID, methodID, updatedPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
