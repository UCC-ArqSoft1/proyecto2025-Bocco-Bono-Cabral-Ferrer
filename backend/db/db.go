package db

import (
	"fmt"
	"gym-api/backend/dao"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLDB struct {
	DB *gorm.DB
}

type Databases interface {
	Connect()
	Migrate()
}

func (db *MySQLDB) Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println(" Error al conectar a la base de datos")
		log.Fatal(err)
		return
	}
	dsn := fmt.Sprintf("%s:%v@tcp(%s:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println(" Error al conectar a la base de datos")
		log.Fatal(err)
	} else {
		log.Println(" Conexión a base de datos exitosa")
	}
}

func (db *MySQLDB) Migrate() {
	err := db.DB.AutoMigrate(
		&dao.UserType{},
		&dao.User{},
		&dao.Activity{},
		&dao.ActivitySchedule{},
		&dao.Enrollment{},
	)

	if err != nil {
		log.Fatal(" Error al migrar modelos:", err)
	} else {
		log.Println("Finishing Migration Database Tables")
	}
}
