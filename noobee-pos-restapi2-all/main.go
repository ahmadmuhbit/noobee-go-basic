package main

import (
	"log"
	"net/http"
	"nobee-pos/database"
	"nobee-pos/employee"
	"nobee-pos/menu"
	"nobee-pos/transaction"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error when load env file with error", err.Error())
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

	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Agent", "Accept"},
	})

	router.Use(c.Handler)

	menu.Register(router, db)
	employee.Register(router, db)
	transaction.Register(router, db)

	const port = ":8000"

	log.Println("Server running at port", port)
	http.ListenAndServe(port, router)
}
