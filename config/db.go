package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB connects to PostgreSQL database using environment variables
func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbName, dbUsername, dbPassword)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = CreateContactTableIfNotExists(DB)
	if err != nil {
		log.Fatal(err)
	}

	err = CreatePaymentTableIfNotExists(DB)
	if err != nil {
		log.Fatal(err)
	}

	err = CreatePaymentMethodTableIfNotExists(DB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

// CreateDatabaseIfNotExists creates a PostgreSQL database if it does not exist
func CreatePaymentDatabaseIfNotExists(db *sql.DB) error {
	dbName := os.Getenv("DB_NAME")

	// Check if the database exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		// Database does not exist, create it
		_, err := db.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			return err
		}
		fmt.Printf("Database %s created\n", dbName)
	} else {
		fmt.Printf("Database %s already exists\n", dbName)
	}

	return nil
}

// CreateTableIfNotExists creates the Payment table if it does not exist
func CreatePaymentTableIfNotExists(db *sql.DB) error {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS payments (
		id SERIAL PRIMARY KEY,
		contact_id INTEGER,
		payment_type VARCHAR(255),
		status VARCHAR(255),
		payer VARCHAR(255),
		billing_order INTEGER,
		condition_related_to VARCHAR(255),
		billing_id VARCHAR(255),
		relationship_to_insured VARCHAR(255),
		insured_full_name VARCHAR(255),
		insured_id VARCHAR(255),
		dob_of_insured DATE,
		address_of_insured VARCHAR(255),
		city_of_insured VARCHAR(255),
		state_of_insured VARCHAR(255),
		zip_code_of_insured VARCHAR(10),
		insurance_start_date DATE,
		insurance_end_date DATE,
		insurance_co_pay DECIMAL(10,2),
		insurance_deductible DECIMAL(10,2),
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	fmt.Println("Table 'payments' created or already exists")
	return nil
}

// CreatePaymentMethodTableIfNotExists creates the PaymentMethods table if it does not exist
func CreatePaymentMethodTableIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS payment_methods (
		id SERIAL PRIMARY KEY,
		contact_id INTEGER,
		method_type VARCHAR(255),
		status VARCHAR(255),
		payer VARCHAR(255),
		billing_order INTEGER,
		condition_related_to VARCHAR(255),
		billing_id VARCHAR(255),
		relationship_to_insured VARCHAR(255),
		insured_full_name VARCHAR(255),
		insured_id VARCHAR(255),
		dob_of_insured DATE,
		address_of_insured VARCHAR(255),
		city_of_insured VARCHAR(255),
		state_of_insured VARCHAR(255),
		zip_code_of_insured VARCHAR(10),
		insurance_start_date DATE,
		insurance_end_date DATE,
		insurance_co_pay DECIMAL(10,2),
		insurance_deductible DECIMAL(10,2),
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (contact_id) REFERENCES contacts(id)
	)`)
	if err != nil {
		return err
	}

	fmt.Println("Table 'payment_methods' created or already exists")
	return nil
}

// CreateContactTableIfNotExists creates the Contacts table if it does not exist
func CreateContactTableIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS contacts (
		id SERIAL PRIMARY KEY,
		contact_type VARCHAR(255),
		email_address VARCHAR(255),
		full_name VARCHAR(255),
		location_id VARCHAR(255),
		owner_user_id VARCHAR(255),
		phone_number VARCHAR(255),
		pipeline_id VARCHAR(255),
		pipeline_stage_id VARCHAR(255),
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	fmt.Println("Table 'contacts' created or already exists")
	return nil
}
