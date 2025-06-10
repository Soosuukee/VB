package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"vb/internal/models"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env non trouvé, je continue avec les variables système.")
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := "5432"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Erreur de connexion à PostgreSQL :", err)
	}

	DB = db
	fmt.Println("✅ Connexion à PostgreSQL réussie")

	// 👉 AutoMigrate à l'intérieur de Connect()
	err = DB.AutoMigrate(&models.Game{}, &models.User{}, &models.Event{})
	if err != nil {
		log.Fatal("❌ Erreur migration :", err)
	}
}
