package handler

import (
	"log"
	"net/http"
	"telegram-clicker-game-be/application"
	"telegram-clicker-game-be/config"

	"github.com/gin-gonic/gin"
)

var (
	rGin *gin.Engine
)

func init() {
	cfg, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	app, err := application.SetupVercel(cfg)
	if err != nil {
		log.Fatal("App: ", err.Error())
	}

	rGin, err = application.SetupGin(app, cfg)
	if err != nil {
		log.Fatal("go-gin: ", err.Error())
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	rGin.ServeHTTP(w, r)
}
