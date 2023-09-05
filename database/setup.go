package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
    // Menggunakan godotenv untuk membaca variabel lingkungan dari file .env
    err := godotenv.Load(".env")
    if err != nil {
        return nil, err
    }

    DB_HOST := os.Getenv("DB_HOST")
    DB_USERNAME := os.Getenv("DB_USERNAME")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_NAME := os.Getenv("DB_NAME")
    DB_PORT := os.Getenv("DB_PORT")

	fmt.Println(DB_HOST, DB_PASSWORD)
    // Membangun string koneksi DSN dengan nilai dari variabel lingkungan
    dsn := "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT

    // Membuka koneksi database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
