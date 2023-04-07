package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"

	///"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

type Setting struct {
	DB_HOST            string
	MYSQL_DB_NAME      string
	MYSQL_DB_USER      string
	MYSQL_DB_PASS      string
	MYSQL_DB_PORT      string
	POSTGRESS_DB_PORT  string
	POSTGGRESS_DB_NAME string
	POSTGRESS_DB_USER  string
	POSTGRESS_DB_PASS  string
	SQLITE_DB_NAME     string
}

func InitilializeSetting() Setting {
	log.Println("Loading .env")
	godotenv.Load()
	DB_HOST := os.Getenv("DB_HOST")
	MYSQL_DB_NAME := os.Getenv("MYSQL_DB_NAME")
	MYSQL_DB_USER := os.Getenv("MYSQL_DB_USER")
	MYSQL_DB_PASS := os.Getenv("MYSQL_DB_PASSWORD")
	MYSQL_DB_PORT := os.Getenv("MYSQL_DB_PORT")
	POSTGRESS_DB_PORT := os.Getenv("POSTGRES_DB_PORT")
	POSTGRES_DB_USER := os.Getenv("POSTGRES_DB_USER")
	POSTGRES_DB_PASSWORD := os.Getenv("POSTGRES_DB_PASSWORD")
	POSTGRES_DB_NAME := os.Getenv("POSTGRES_DB_NAME")
	DB_TYPE := os.Getenv("DB_TYPE")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	SQLITE_DB_NAME := os.Getenv("SQLITE_DB_NAME")

	switch {
	case DB_HOST == "":
		log.Println("DB HOST Non Set")
		os.Exit(1)
	case DB_TYPE == "":
		log.Println("DB_TYPE Non set")
		os.Exit(1)
	case MYSQL_DB_NAME == "":
		log.Println("MYSQL_DB_NAME Non Set")
		os.Exit(1)
	case MYSQL_DB_USER == "":
		log.Println("MYSQL_DB_USER Non Set")
		os.Exit(1)
	case MYSQL_DB_PASS == "":
		log.Println("MYSQL_DB_PASS Non Set")
		os.Exit(1)
	case POSTGRESS_DB_PORT == "":
		log.Println("POSTGRESS_DB_PORT Non Set")
		os.Exit(1)
	case POSTGRES_DB_NAME == "":
		log.Println("POSTGRESS_DB_NAME Non Set")
		os.Exit(1)
	case POSTGRES_DB_USER == "":
		log.Println("POSTGRESS_DB_USER Non Set")
		os.Exit(1)
	case POSTGRES_DB_PASSWORD == "":
		log.Println("POSTGRESS_DB_PASSWORD Non Set")
		os.Exit(1)
	case SQLITE_DB_NAME == "":
		log.Println("SQLITE_DB_NAME Non Set")
		os.Exit(1)
	case SERVER_PORT == "":
		log.Println("SERVER_PORT Non Set")
		os.Exit(1)
	case MYSQL_DB_PORT == "":
		log.Println("MYSQL_DB_PORT Non Set")
		os.Exit(1)
	}
	setting := Setting{
		SQLITE_DB_NAME:     SQLITE_DB_NAME,
		MYSQL_DB_NAME:      MYSQL_DB_NAME,
		MYSQL_DB_USER:      MYSQL_DB_USER,
		MYSQL_DB_PASS:      MYSQL_DB_PASS,
		MYSQL_DB_PORT:      MYSQL_DB_PORT,
		DB_HOST:            DB_HOST,
		POSTGRESS_DB_PORT:  POSTGRESS_DB_PORT,
		POSTGGRESS_DB_NAME: POSTGRES_DB_NAME,
		POSTGRESS_DB_USER:  POSTGRES_DB_USER,
		POSTGRESS_DB_PASS:  POSTGRES_DB_PASSWORD,
	}
	return setting
}

func Connect() {
	data := InitilializeSetting()
	d, err := gorm.Open(sqlite.Open(data.SQLITE_DB_NAME), &gorm.Config{})
	if err != nil {
		log.Println("Error in connecting", err)
	}
	db = d
	log.Println("Databse")
}

func GetDb() *gorm.DB {
	return db
}
