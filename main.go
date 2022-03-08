package main

import (
	"api/config"
	"api/database"
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	database.StartDB(config.BaseUrl)

	r := router.NewRouter()

	fmt.Println("Rodando API!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
