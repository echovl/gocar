package main

import (
	"log"
	"os"

	"github.com/echovl/gocar/http"
	"github.com/echovl/gocar/postgres"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("POSTGRES_DSN")
	port := os.Getenv("PORT")

	carStg, err := postgres.NewCarStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}

	sv := http.NewServer(carStg)

	err = sv.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
