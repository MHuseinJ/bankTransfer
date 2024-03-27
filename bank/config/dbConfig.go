package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DBConfig struct {
	DB *gorm.DB
}

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	DBPort   string
	SSLMode  string
	TimeZone string
}

func setupDB(config *Config) (*gorm.DB, error) {
	//setup connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.DBPort,
		config.SSLMode,
		config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func LoadDB() DBConfig {
	configDB := Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		DBPort:   os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
	db, err := setupDB(&configDB)
	if err != nil {
		log.Fatal(err)
	}
	dbConfig := DBConfig{
		DB: db,
	}
	return dbConfig
}

func (dc DBConfig) MigrateTable(tables ...interface{}) {
	err := dc.DB.AutoMigrate(tables)
	if err != nil {
		log.Fatal(err)
	}
}
