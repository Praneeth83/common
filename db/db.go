package common // important: package name must match the module import

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global DB variable
var DB *gorm.DB

func Connect() {
	dsn := "postgres://postgres:praneeth@localhost:5432/mydb?sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	fmt.Println("Connected to remote DB")
}
