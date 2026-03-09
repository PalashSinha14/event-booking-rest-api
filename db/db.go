package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	env := os.Getenv("ENV")

	envFile := ".env.local"
	if env == "docker" {
		envFile = ".env.docker"
	}

	// Try loading env file but DO NOT crash if missing
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("%s not found. Using system environment variables.\n", envFile)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		getSSLMode(),
	)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to create DB instance:", err)
	}

	// Retry DB connection
	maxRetries := 90
	retryInterval := 2 * time.Second

	for i := 0; i < maxRetries; i++ {

		err = DB.Ping()
		if err == nil {
			log.Println("Connected to PostgreSQL!")
			break
		}

		log.Printf("PostgreSQL not ready yet... Retrying (%d/%d)\n", i+1, maxRetries)
		time.Sleep(retryInterval)
	}

	if err != nil {
		log.Fatal("PostgreSQL database not reachable after retries.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	createTables()
}

func getSSLMode() string {

	ssl := os.Getenv("DB_SSLMODE")

	if ssl == "" {
		// Local default
		return "disable"
	}

	return ssl
}





/*
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	env := os.Getenv("ENV")

	envFile := ".env.local"
	if env == "docker" {
		envFile = ".env.docker"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Open connection (does NOT actually connect yet)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to create DB instance:", err)
	}

	// Retry logic — up to 3 minutes total
	maxRetries := 90 // 90 × 2s = 180 seconds (3 minutes)
	retryInterval := 2 * time.Second

	for i := 0; i < maxRetries; i++ {

		err = DB.Ping()
		if err == nil {
			log.Println("Connected to PostgreSQL!")
			break
		}

		log.Printf("PostgreSQL not ready yet... Retrying (%d/%d)\n", i+1, maxRetries)
		time.Sleep(retryInterval)
	}

	if err != nil {
		log.Fatal("PostgreSQL database not reachable after retries.")
	}

	// Adjusted for 4GB RAM system
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	createTables()
}*/

/*
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	env := os.Getenv("ENV")

	envFile := ".env.local"
	if env == "docker" {
		envFile = ".env.docker"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var dbInstance *sql.DB

	for i := 0; i < 10; i++ {

		dbInstance, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Println("Waiting for PostgreSQL to start...")
			time.Sleep(2 * time.Second)
			continue
		}

		err = dbInstance.Ping()
		if err == nil {
			log.Println("Connected to PostgreSQL!")
			DB = dbInstance
			break
		}

		log.Println("PostgreSQL not ready yet... Retrying in 2s")
		time.Sleep(2 * time.Second)
	}

	if DB == nil {
		log.Fatal("PostgreSQL database not reachable after retries.")
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)

	createTables()
}*/

/*
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	env := os.Getenv("ENV")

	envFile := ".env.local"
	if env == "docker" {
    	envFile = ".env.docker"
	}

	err := godotenv.Load(envFile)
	if err != nil {
    	log.Fatalf("Error loading %s file", envFile)
	}


//	err := godotenv.Load()
//	if err != nil {
//		panic("Error loading .env file")
//	}


	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	//var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Could not connect to PostgreSQL database.")
	}

	err = DB.Ping()
	if err != nil {
		panic("PostgreSQL database not reachable.")
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)

	createTables()

}
*/

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime TIMESTAMP NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table")
	}
}

/*
package db

import (
	"database/sql"
	_ "modernc.org/sqlite"

)

var DB *sql.DB

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite","api.db")
	if err != nil{
		panic("Could not connect to database.")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables(){

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil{
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES user(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil{
		panic("Could not create events table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil{
		panic("Could not create registrations table.")
	}
}*/
