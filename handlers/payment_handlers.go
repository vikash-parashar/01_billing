package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/vikash-parashar/01_billing/controllers"
	"github.com/vikash-parashar/01_billing/models"
)

// CreatePaymentMethodHandler handles the creation of a new payment method for a given contact ID.
func CreatePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
	contactID := chi.URLParam(r, "id")

	// Parse the request body to get the payment method details
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Convert payment type and relationship to insured to lowercase for case-insensitive checks
	payment.PaymentType = strings.ToLower(payment.PaymentType)
	payment.RelationshipToInsured = strings.ToLower(payment.RelationshipToInsured)

	if payment.PaymentType == "cash" {
		if payment.RelationshipToInsured == "self" {
			contactID, _ := strconv.Atoi(contactID)
			insured := &models.Insured{
				ContactID:   contactID,
				FullName:    payment.InsuredFullName,
				DateOfBirth: payment.DOBOfInsured,
				Address:     payment.AddressOfInsured,
				City:        payment.CityOfInsured,
				State:       payment.StateOfInsured,
				ZipCode:     payment.ZipCodeOfInsured,
			}

			if id, err := controllers.CreateInsured(*insured); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				log.Printf("New insured created with ID: %d\n", id)
			}
		}
	}

	// Validate and handle the payment method creation
	if res, err := controllers.CreatePaymentMethod(contactID, payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if id, err := res.LastInsertId(); err != nil {
			log.Println(err)
		} else {
			log.Printf("New payment inserted with ID: %d\n", id)
		}
	}

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
