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

	err = createTables(DB)
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

func createTables(db *sql.DB) error {
	// Create Contact table
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Contact (
            ID SERIAL PRIMARY KEY,
            ContactType VARCHAR(50),
            EmailAddress VARCHAR(255),
            FullName VARCHAR(255),
            LocationID VARCHAR(50),
            OwnerUserID VARCHAR(50),
            PhoneNumber VARCHAR(20),
            PipelineID VARCHAR(50),
            PipelineStageID VARCHAR(50),
            CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	// Create Insured table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Insured (
            ID SERIAL PRIMARY KEY,
            ContactID INT REFERENCES Contact(ID) ON DELETE CASCADE,
            FullName VARCHAR(255),
            DateOfBirth VARCHAR(20),
            Address VARCHAR(255),
            City VARCHAR(50),
            Country VARCHAR(50),
            State VARCHAR(50),
            ZipCode VARCHAR(20)
        )
    `)
	if err != nil {
		return err
	}

	// Create Payment table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Payment (
            ID SERIAL PRIMARY KEY,
            ContactID INT REFERENCES Contact(ID) ON DELETE CASCADE,
            PaymentType VARCHAR(50),
            Status VARCHAR(50),
            Payer VARCHAR(255),
            BillingOrder INT,
            ConditionRelatedTo VARCHAR(255),
            BillingID VARCHAR(50),
            RelationshipToInsured VARCHAR(20),
            InsuredFullName VARCHAR(255),
            InsuredID VARCHAR(50),
            DOBOfInsured VARCHAR(20),
            AddressOfInsured VARCHAR(255),
            CityOfInsured VARCHAR(50),
            StateOfInsured VARCHAR(50),
            ZipCodeOfInsured VARCHAR(20),
            InsuranceStartDate VARCHAR(20),
            InsuranceEndDate VARCHAR(20),
            InsuranceCoPay DECIMAL(10, 2),
            InsuranceDeductible DECIMAL(10, 2),
            CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	return nil
}
