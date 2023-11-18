package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/database"
	"rest-api/menu"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error when load .env file with error", err.Error())
	}

	db, err := database.ConnectPostgres(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db not connected")
	}

	router := chi.NewRouter()

	menu.Register(router, db)
	const port = ":8000"

	log.Println("Server running at port", port)
	http.ListenAndServe(port, router)
}
