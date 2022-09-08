package main

import (
	"github.com/gin-gonic/gin"
)

type application struct {
	AppName string
	Server  Server
	Debug   bool
}

type Server struct {
	Host string
	Port string
	Url  string
}

func main() {
	app := application{
		AppName: "Dokander",
		Server: Server{
			Host: "localhost",
			Port: "8080",
			Url:  "http://localhost:8080",
		},
		Debug: true,
	}
	if !app.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	app.ListenAndServe()
}
