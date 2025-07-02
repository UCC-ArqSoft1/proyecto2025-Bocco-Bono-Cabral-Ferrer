package db

import (
	"fmt"
	"gym-api/dao"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct {
	DB *gorm.DB
}

type Databases interface {
	Connect()
	Migrate()
	WaitForDB(dsn string) (*gorm.DB, error)
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

	db.DB, err = db.waitForDB(dsn)

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

func (db *MySQLDB) waitForDB(dsn string) (*gorm.DB, error) {
	var err error
	var gormDB *gorm.DB

	for i := 0; i < 10; i++ {
		gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := gormDB.DB()
			if err == nil && sqlDB.Ping() == nil {
				return gormDB, nil
			}
		}
		log.Printf("⏳ Esperando base de datos... intento %d/10\n", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("❌ no se pudo conectar a la BD: %v", err)
}
