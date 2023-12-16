package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/vikash-parashar/01_billing/config"
	"github.com/vikash-parashar/01_billing/handlers"
)

func LoadEnv() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// func GetEnv(key string) string {
// 	value, exists := os.LookupEnv(key)
// 	if !exists {
// 		log.Fatalf("Environment variable %s not set", key)
// 	}
// 	return value
// }

// Initialize the environment variables
func init() {
	LoadEnv()
}

// @title Billing API
// @version 1.0
// @description API for managing billing and payments.
// @host localhost:8080
// @BasePath /v1
func main() {
	appPort := os.Getenv("APP_PORT")

	// Connect to the database
	config.ConnectDB()

	// Initialize the Chi router
	r := chi.NewRouter()

	r.Get("/contact", handlers.GetAllContact)
	r.Post("/contact", handlers.CreateNewContact)
	r.Get("/contact/{contactID}", handlers.GetContactByContactID)
	r.Put("/contact/{contactID}", handlers.UpdateContactByContactID)
	r.Delete("/contact/{contactID}", handlers.DeleteContactByContactID)

	r.Get("/insured", handlers.GetAllInsured)
	r.Post("/insured", handlers.CreateNewInsured)
	r.Get("/insured/{insuredID}", handlers.GetInsuredByInsuredID)
	r.Put("/insured/{insuredID}", handlers.UpdateInsuredByInsuredID)
	r.Delete("/insured/{insuredID}", handlers.DeleteInsuredByInsuredID)

	r.Get("/payment", handlers.GetAllPayments)
	r.Post("/payment", handlers.CreateNewPayment)
	r.Get("/payment/{paymentID}", handlers.GetPaymentByPaymentID)
	r.Put("/payment/{paymentID}", handlers.UpdatePaymentByPaymentID)
	r.Delete("/payment/{paymentID}", handlers.DeletePaymentByPaymentID)

	// All CRUD operations for contact
	r.Get("/contacts/{contactID}/insured", handlers.GetAllInsuredByContactID)
	r.Post("/contacts/{contactID}/insured", handlers.CreateNewInsuredByContactID)
	r.Get("/contacts/{contactID}/insured/{insuredID}", handlers.GetInsuredByContactAndInsuredID)
	r.Put("/contacts/{contactID}/insured/{insuredDd}", handlers.UpdateInsuredByContactAndInsuredID)
	r.Delete("/contacts/{contactID}/insured/{insuredID}", handlers.DeleteInsuredByContactAndInsuredID)

	// CRUD endpoints for contacts payment
	r.Post("/contact/{contactID}/payment", handlers.CreateNewPaymentByContactID)
	r.Put("/contact/{contactID}/payment/{paymentID}", handlers.UpdatePaymentByContactAndPaymentID)
	r.Delete("/contact/{contactID}/payment/{paymentID}", handlers.DeletePaymentByContactAndPaymentID)
	r.Get("/contact/{contactID}/payment/{paymentID}", handlers.GetPaymentByContactAndPaymentID)
	r.Get("/contact/{contactID}/payment", handlers.GetAllPaymentsByContactID)

	// CRUD endpoints for insured payment
	r.Post("/insured/{insuredID}/payment", handlers.CreateNewPaymentByInsuredID)
	r.Put("/insured/{insuredID}/payment/{paymentID}", handlers.UpdatePaymentByPaymentAndInsuredID)
	r.Delete("/insured/{insuredID}/payment/{paymentID}", handlers.DeletePaymentByPaymentAndInsuredID)
	r.Get("/insured/{insuredID}/payment/{paymentID}", handlers.GetPaymentByPaymentAndInsuredID)
	r.Get("/insured/{insuredID}/payment", handlers.GetAllPaymentsByInsuredID)

	// Run the application
	log.Fatalln(http.ListenAndServe(":"+appPort, r))
}
