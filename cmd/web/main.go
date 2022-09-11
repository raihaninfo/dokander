package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raihaninfo/dokander/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type application struct {
	AppName string
	Server  Server
	Debug   bool
	db      *gorm.DB
}

type Server struct {
	Host string
	Port string
	Url  string
}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	// db.Statement.SQL.Reset()
	// db.Commit().Statement.SQL.Reset()
	// db.Debug().Statement.SQL.Reset()
	db.AutoMigrate(&models.Products{})

	app := application{
		AppName: "Dokander",
		Server: Server{
			Host: "localhost",
			Port: "8080",
			Url:  "http://localhost:8080",
		},
		Debug: true,
		db:    db,
	}
	if !app.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	app.ListenAndServe()
}

func openDB() (*gorm.DB, error) {
	psqInfo := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "dev", "secret", "dokander")

	db, err := gorm.Open(postgres.Open(psqInfo), &gorm.Config{})
	return db, err
}
