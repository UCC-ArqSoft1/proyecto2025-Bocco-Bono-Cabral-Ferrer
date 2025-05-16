package db

import (
	"log"

	"gym-api/backend/clients"
	"gym-api/backend/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Println(" Error al conectar a la base de datos")
		log.Fatal(err)
		return
	}
	// Conexión a base de datos SQLite en memoria (solo para pruebas)
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// Para usar archivo físico:
	// DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Println(" Error al conectar a la base de datos")
		log.Fatal(err)
	} else {
		log.Println(" Conexión a base de datos exitosa")
	}

	// AutoMigrate: crea las tablas si no existen
	err = DB.AutoMigrate(
		&domain.UserType{},
		&domain.User{},
		&domain.Activity{},
		&domain.Enrollment{},
	)

	if err != nil {
		log.Fatal(" Error al migrar modelos:", err)
	} else {
		log.Println("Finishing Migration Database Tables")
	}

	// Inyectar DB en los paquetes clients
	clients.UserClient = DB
	clients.ActivityClient = DB
	clients.UserTypeClient = DB
	clients.EnrollmentClient = DB
}
