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

	r.Post("/contacts/{id}/payment-methods", handlers.CreatePaymentMethodHandler)
	r.Put("/contacts/{id}/payment-methods/{methodId}", handlers.UpdatePaymentMethodHandler)
	r.Delete("/contacts/{id}/payment-methods/{methodId}", handlers.DeletePaymentMethodHandler)
	r.Get("/contacts/{id}/payment-methods/{methodId}", handlers.GetPaymentMethodHandlerByMethodID)
	r.Get("/contacts/{id}/payment-methods", handlers.GetAllPaymentMethodsForContactHandler)

	// Run the application
	log.Fatalln(http.ListenAndServe(":"+appPort, r))
}
