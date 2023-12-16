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

func GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := controllers.GetAllPayments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
}

func CreateNewPayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPayment, err := controllers.CreatePayment(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPayment)
}

func GetPaymentByPaymentID(w http.ResponseWriter, r *http.Request) {
	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	payment, err := controllers.GetPaymentByID(paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

func UpdatePaymentByPaymentID(w http.ResponseWriter, r *http.Request) {
	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	var updatedPayment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&updatedPayment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := controllers.UpdatePaymentByID(paymentID, updatedPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

func DeletePaymentByPaymentID(w http.ResponseWriter, r *http.Request) {
	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	err = controllers.DeletePaymentByID(paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllPaymentsByContactID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	payments, err := controllers.GetAllPaymentsByContactID(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payments)
}

func CreateNewPaymentByContactID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPayment, err := controllers.CreatePaymentByContactID(contactID, payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPayment)
}

func GetPaymentByContactAndPaymentID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	payment, err := controllers.GetPaymentByContactAndPaymentID(contactID, paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

func UpdatePaymentByContactAndPaymentID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	var updatedPayment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&updatedPayment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := controllers.UpdatePaymentByContactAndPaymentID(contactID, paymentID, updatedPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

func DeletePaymentByContactAndPaymentID(w http.ResponseWriter, r *http.Request) {
	contactIDStr := chi.URLParam(r, "contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	err = controllers.DeletePaymentByContactAndPaymentID(contactID, paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllPaymentsByInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	payments, err := controllers.GetAllPaymentsByInsuredID(insuredID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payments)
}

func CreateNewPaymentByInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPayment, err := controllers.CreatePaymentByInsuredID(insuredID, payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPayment)
}

func GetPaymentByPaymentAndInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	payment, err := controllers.GetPaymentByPaymentAndInsuredID(insuredID, paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

func UpdatePaymentByPaymentAndInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	var updatedPayment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&updatedPayment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := controllers.UpdatePaymentByPaymentAndInsuredID(insuredID, paymentID, updatedPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

func DeletePaymentByPaymentAndInsuredID(w http.ResponseWriter, r *http.Request) {
	insuredIDStr := chi.URLParam(r, "insuredID")
	insuredID, err := strconv.Atoi(insuredIDStr)
	if err != nil {
		http.Error(w, "Invalid insured ID", http.StatusBadRequest)
		return
	}

	paymentIDStr := chi.URLParam(r, "paymentID")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	err = controllers.DeletePaymentByPaymentAndInsuredID(insuredID, paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
