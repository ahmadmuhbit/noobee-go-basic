package main

import (
	"log"
	"net/http"
	"nobee-pos/auth"
	"nobee-pos/database"
	routerChi "nobee-pos/infra/router/chi"
	"nobee-pos/menu"
	"nobee-pos/utility"
	"os"

	"github.com/go-chi/chi/v5"
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

	utility.InitToken("INI_SECRET", 1*60)

	router := chi.NewRouter()

	router.Use(routerChi.Logger, routerChi.CORS)

	menu.Register(router, db)
	auth.Register(router, db)

	const port = ":8000"

	log.Println("Server running at port", port)
	http.ListenAndServe(port, router)
}
