package main

import (
  "database/sql"
  "fmt"
  "log"
  "os"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "github.com/joho/godotenv"
  _ "github.com/jackc/pgx/v5/stdlib"
)

func main () {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// build dsn
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	// open sqp connection using pgx driver
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// wrap sqldb in gorm
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
	Conn: sqlDB,
	}), &gorm.Config{})
	_ = gormDB 

	if err != nil {
		log.Fatal("Failed to initialize gorm:", err)
	}

	fmt.Println("Successfully connected to the database using GORM with pgx driver!")
}