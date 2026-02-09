package main

import (
  "database/sql"
  "fmt"
  "log"
  "os"
  "time"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "github.com/joho/godotenv"
  _ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID  uint  `gorm:"primaryKey"`
	Name  string  `gorm:"size:100; not null"`
	Email string  `gorm:"size:150; unique; not null"`
	Phone string  `gorm:"size:20"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

	// _ = gormDB 

	if err != nil {
		log.Fatal("Failed to initialize gorm:", err)
	}
	fmt.Println("Successfully connected to the database using GORM with pgx driver!")

	err = gormDB.AutoMigrate(&User{})

	if err != nil {
		log.Fatal("Failed to auto-migrate User model:", err)
	}

	fmt.Println("user table auto-migrated successfully!")

	// create a record using create method
	user := User{
		Name: "Sadia",
		Email: "sadia@example.com",
		Phone: "+8801727392836",
	}

	result := gormDB.Create(&user)

	if result.Error != nil {
		log.Fatal("Failed to create user:", result.Error)
	}
	fmt.Printf("User created successfully with ID: %v\n", user.ID)

	// fetch users using raw sql query
	var users []User
	gormDB.Raw("select * from users").Scan(&users)

	for _, u := range users {
		fmt.Printf("ID: %v, Name: %v, Email: %v, Phone: %v\n", u.ID, u.Name, u.Email, u.Phone)
	}

	// fetch user by id
	var user1 User
	gormDB.Raw("select * from users where id = ?", 1).Scan(&user1)
	fmt.Printf("User with ID 1: Name: %v, Email: %v, Phone: %v\n", user1.Name, user1.Email, user1.Phone)
}