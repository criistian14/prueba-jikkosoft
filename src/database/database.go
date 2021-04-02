package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

// * Database struct
type Database struct {
	DB *gorm.DB

	migrate    bool
	connection string
	host       string
	port       string
	name       string
	user       string
	password   string
}

// * Get status to migrate
func (db *Database) GetStatusMigrate() bool {
	err := db.getVars()
	if err != nil {
		return false
	}

	return db.migrate
}

func (db *Database) getVars() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	db.connection = os.Getenv("DB_CONNECTION")
	db.host = os.Getenv("DB_HOST")
	db.port = os.Getenv("DB_PORT")
	db.name = os.Getenv("DB_DATABASE")
	db.user = os.Getenv("DB_USERNAME")
	db.password = os.Getenv("DB_PASSWORD")
	db.migrate, _ = strconv.ParseBool(os.Getenv("DB_MIGRATE"))

	return nil
}

// * New database connexion and return
func NewDBClient() (Database, error) {
	var dbTemp *gorm.DB
	var err error

	var client Database
	err = client.getVars()

	configDB := &gorm.Config{}

	// Establish connection according to the database
	switch client.connection {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", client.user, client.password, client.host, client.name)
		dbTemp, err = gorm.Open(mysql.Open(dsn), configDB)
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", client.host, client.port, client.user, client.name, client.password)
		dbTemp, err = gorm.Open(postgres.Open(dsn), configDB)
	}

	if err != nil {
		return client, err
	}

	client.DB = dbTemp
	return client, nil
}

// * Close data base to performance
func (db *Database) CloseDataBase() {
	client, _ := db.DB.DB()
	_ = client.Close()
}
