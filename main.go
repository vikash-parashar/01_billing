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
	r.Get("/insured", handlers.GetAllInsuredHandler)
	r.Post("/insured", handlers.CreateInsuredHandler)
	r.Get("/insured/{id}", handlers.GetInsuredByIDHandler)
	r.Put("/insured/{id}", handlers.UpdateInsuredByIDHandler)
	r.Delete("/insured/{id}", handlers.DeleteInsuredByIDHandler)

	r.Get("/contact", handlers.GetAllContactHandler)
	r.Post("/contact", handlers.CreateContactHandler)
	r.Get("/contact/{id}", handlers.GetContactByIDHandler)
	r.Put("/contact/{id}", handlers.UpdateContactByIDHandler)
	r.Delete("/contact/{id}", handlers.DeleteContactByIDHandler)

	// All CRUD operations for contact
	r.Get("/contacts/{id}/insured", handlers.GetAllInsuredHandler)
	r.Post("/contacts/{id}/insured", handlers.CreateInsuredHandler)
	r.Get("/contacts/{id}/insured/{id}", handlers.GetInsuredByIDHandler)
	r.Put("/contacts/{id}/insured/{id}", handlers.UpdateInsuredByIDHandler)
	r.Delete("/contacts/{id}/insured/{id}", handlers.DeleteInsuredByIDHandler)

	// CRUD endpoints for contacts payment
	r.Post("/contacts/{id}/payment-methods", handlers.CreatePaymentMethodHandler)
	r.Put("/contacts/{id}/payment-methods/{methodId}", handlers.UpdatePaymentMethodHandler)
	r.Delete("/contacts/{id}/payment-methods/{methodId}", handlers.DeletePaymentMethodHandler)
	r.Get("/contacts/{id}/payment-methods/{methodId}", handlers.GetPaymentMethodHandlerByMethodID)
	r.Get("/contacts/{id}/payment-methods", handlers.GetAllPaymentMethodsForContactHandler)

	// CRUD endpoints for insured payment
	// r.Post("/insured/{id}/payment-methods", handlers.CreatePaymentMethodForInsuredHandler)
	// r.Put("/insured/{id}/payment-methods/{methodId}", handlers.UpdatePaymentMethodForInsuredHandler)
	// r.Delete("/insured/{id}/payment-methods/{methodId}", handlers.DeletePaymentMethodForInsuredHandler)
	// r.Get("/insured/{id}/payment-methods/{methodId}", handlers.GetPaymentMethodForInsuredHandlerByMethodID)
	// r.Get("/insured/{id}/payment-methods", handlers.GetAllPaymentMethodsForInsuredHandler)

	// Run the application
	log.Fatalln(http.ListenAndServe(":"+appPort, r))
}
