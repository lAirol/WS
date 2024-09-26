package main

import (
	"WS/internal/config/constants"
	"WS/internal/config/site_config"
	"WS/internal/db"
	"WS/internal/extentions/system"
	"WS/internal/route"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var err error
	//загрузка констант
	constants.Constants, err = constants.LoadSiteConstants("./internal/config/constants/constants.json")
	if err != nil {
		log.Fatalf("Error loading constants: %v", err)
	}
	//загрузка конфига
	site_config.SiteConfig, err = site_config.LoadConfig("./internal/config/site_config/config.json")
	if err != nil {
		log.Fatalf("Error loading site config: %v", err)
	}
	SD := site_config.SiteConfig.Database

	// Инициализация базы данных
	db.InitDB(SD.Username, SD.Password, SD.DBName)
	// Маршрутизация
	route.InitRoute()
	//Мониторинг
	go system.InitSystemWatch(constants.Constants.SystemWatchInterval * time.Second)

	fmt.Println("Сервер запущен на http://localhost:8080/")
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		panic(err)
	}
}
