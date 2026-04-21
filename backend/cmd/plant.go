package cmd

import (
	"fmt"
	"os"

	"github.com/haley-marie/greenhousedashboard/backend/models"
	"github.com/haley-marie/greenhousedashboard/backend/repository"
)

func HandleAddPlant(repo *repository.PlantRepository) {
	if len(os.Args) < 3 {
		fmt.Println("usage: gh add-plant <species> [nickname] [variety]")
		return
	}

	species := os.Args[2]

	var nickname, variety string

	if len(os.Args) > 3 {
		nickname = os.Args[3]
	}

	if len(os.Args) > 4 {
		variety = os.Args[4]
	}

	p := &models.Plant{
		Species:  species,
		Nickname: nickname,
		Variety:  variety,
	}

	plant, err := repo.AddPlant(p)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("added plant: %s (ID: %d)\n", plant.Species, plant.ID)

}

func HandleListPlants(repo *repository.PlantRepository) {
	plants, err := repo.ListPlants()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if len(plants) == 0 {
		fmt.Println("no plants found")
		return
	}

	for _, plant := range plants {
		if plant.Nickname != "" && plant.Variety != "" {
			fmt.Printf("ID: %d | %s (%s - %s)\n", plant.ID, plant.Nickname, plant.Species, plant.Variety)
		} else if plant.Nickname != "" {
			fmt.Printf("ID: %d | %s (%s)\n", plant.ID, plant.Nickname, plant.Species)
		} else {
			fmt.Printf("ID: %d | %s\n", plant.ID, plant.Species)
		}

	}
}
