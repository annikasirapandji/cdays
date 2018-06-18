package main

import (
	"log"
	"net/http"
	"os"

	"github.com/annikasirapandji/cdays/internal/routing"
	"github.com/annikasirapandji/cdays/internal/version"
)

func main() {
	log.Print("The application is starting, version is %s, build time is %s, commit is %v ...", version.Release, version.BuildTime, version.Commit)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}

	r := routing.NewBLRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
