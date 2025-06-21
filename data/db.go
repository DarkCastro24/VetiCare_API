package data

import (
	"errors"
	"fmt"
	"log"
	"os"

	"PetVet/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func RunPostgresDB() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	if user == "" || password == "" || dbname == "" || host == "" || port == "" {
		return fmt.Errorf("las variables de entorno para la base de datos no están completamente configuradas")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	DB = db
	err = db.AutoMigrate(
		&entities.User{},
		&entities.Admin{},
		&entities.Pet{},
		&entities.Appointment{},
		&entities.AdminType{},
		&entities.UserRole{},
		&entities.Species{},
	)
	if err != nil {
		log.Fatal("Error al migrar las tablas a PostgreSQL: ", err)
	}
	defaultAdminTypes := []entities.AdminType{
		{ID: 1, Type: "Root"},
		{ID: 2, Type: "Admin"},
	}
	for _, at := range defaultAdminTypes {
		var existing entities.AdminType
		result := db.First(&existing, "id = ?", at.ID)
		if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if err := db.Create(&at).Error; err != nil {
				log.Printf("Error insertando AdminType %v: %v\n", at, err)
			} else {
				log.Printf("INSERT INTO AdminType %v\n", at)
			}
		}
	}
	defaultUserRoles := []entities.UserRole{
		{ID: 1, Role: "Dueño"},
		{ID: 2, Role: "Veterinario"},
	}

	for _, ur := range defaultUserRoles {
		var existing entities.UserRole
		result := db.First(&existing, "id = ?", ur.ID)
		if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if err := db.Create(&ur).Error; err != nil {
				log.Printf("Error insertando UserRole %v: %v\n", ur, err)
			} else {
				log.Printf("INSERT INTO UserRole %v\n", ur)
			}
		}
	}
	defaultSpecies := []entities.Species{
		{ID: 1, Name: "Perro", ImageURL: "/images/species/perro.png"},
		{ID: 2, Name: "Gato", ImageURL: "/images/species/gato.png"},
		{ID: 3, Name: "Ave", ImageURL: "/images/species/ave.png"},
	}
	for _, sp := range defaultSpecies {
		var existing entities.Species
		result := db.First(&existing, "id = ?", sp.ID)
		if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if err := db.Create(&sp).Error; err != nil {
				log.Printf("Error insertando Species %v: %v\n", sp, err)
			} else {
				log.Printf("INSERT INTO Species  %v\n", sp)
			}
		}
	}
	return nil
}
