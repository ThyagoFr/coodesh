package main

import (
	"github.com/joho/godotenv"
	"github.com/thyagofr/coodesh/desafio/api"
	"github.com/thyagofr/coodesh/desafio/database"
	"github.com/thyagofr/coodesh/desafio/utils"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Arquivo .env nao encontrado")
	}
}

func main() {
	server := api.Routes()
	database.InitDatabase()
	utils.InitializeCron()
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
