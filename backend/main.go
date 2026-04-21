package main

import (
	"log"
	"os"

	"github.com/haley-marie/greenhousedashboard/backend/cmd"
	"github.com/haley-marie/greenhousedashboard/backend/db"
	"github.com/haley-marie/greenhousedashboard/backend/repository"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("expected command")
		return
	}

	database, err := db.OpenDatabase("../data/greenhouse.db")
	if err != nil {
		log.Fatalf("error opening database: %s...exiting", err)
	}
	defer database.Close()

	err = db.RunMigrations(database)
	if err != nil {
		log.Fatalf("error running migrations: %s...exiting", err)
	}

	repo := repository.PlantRepository{DB: database}

	command := os.Args[1]

	switch command {
	case "add-plant":
		cmd.HandleAddPlant(&repo)
	case "list-plants":
		cmd.HandleListPlants(&repo)
	case "get-plant":
		cmd.HandleGetPlantByID(&repo)
	case "delete-plant":
		cmd.HandleDeletePlantByID(&repo)
	}
}
