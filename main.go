package main

import (
	"fmt"
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Attendance struct {
	ID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID  string `gorm:"type:uuid"`
	InDate  *time.Time
	OutDate *time.Time
}

func StartConnection() *gorm.DB {
	dbHost := "localhost"
	dbPort := "5433"
	dbUser := "monggo"
	dbPass := "monggo"
	dbName := "testing"

	// jika menggunakan heroku maka sslmode harus require (sslmode=require), jika tidak maka sslmode=disable
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to database")
		return nil
	}
	fmt.Println("Success to connect to database")

	db.AutoMigrate(&Attendance{})
	return db
}

func main() {
	db := StartConnection()
	a := Attendance{}

	a.UserID = "1eae463b-471a-4eb2-9911-11cc7cad4d32"
	now := time.Now()
	a.InDate = &now
	db.Create(&a)

	// Mengambil latest date
	// db.Order("in_date DESC").First(&a)
	// fmt.Println(a)

	// Mengambil satu latest date by date
	db.Raw("SELECT * FROM attendances WHERE in_date BETWEEN ? AND ? ORDER BY in_date DESC LIMIT 1", "2022-06-30 00:00:00", "2022-06-30 24:00:00").Scan(&a)
	// fmt.Println(a)
}