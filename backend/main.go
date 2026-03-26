package main

import (
	"log"

	"github.com/haley-marie/greenhousedashboard/backend/db"
)

func main() {
	database, err := db.OpenDatabase("../data/greenhouse.db")
	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}
	defer database.Close()

	err = db.RunMigrations(database)
	if err != nil {
		log.Fatalf("error running migrations: %s", err)
	}

	log.Println("migrations complete")
}
