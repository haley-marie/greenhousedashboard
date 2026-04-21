package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/haley-marie/greenhousedashboard/backend/models"
	"github.com/haley-marie/greenhousedashboard/backend/repository"
)

func HandleAddPlant(repo *repository.PlantRepository) {
	if len(os.Args) < 3 {
		fmt.Println("usage: greenhouse add-plant <species> [nickname] [variety]")
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
		switch {
		case plant.Nickname != "" && plant.Variety != "":
			fmt.Printf("ID: %d | %s (%s - %s)\n", plant.ID, plant.Nickname, plant.Species, plant.Variety)

		case plant.Nickname != "":
			fmt.Printf("ID: %d | %s (%s)\n", plant.ID, plant.Nickname, plant.Species)

		default:
			fmt.Printf("ID: %d | %s\n", plant.ID, plant.Species)
		}
	}
}

func HandleGetPlantByID(repo *repository.PlantRepository) {
	if len(os.Args) < 3 {
		fmt.Println("usage: greenhouse get-plant <id>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	p, err := repo.GetPlantByID(id)
	if err != nil {
		if err == repository.ErrNotFound {
			fmt.Println("plant not found")
		} else {
			fmt.Println("error:", err)
		}
		return
	}

	switch {
	case p.Nickname != "" && p.Variety != "":
		fmt.Printf("ID: %d | %s (%s - %s)\n", p.ID, p.Nickname, p.Species, p.Variety)

	case p.Nickname != "":
		fmt.Printf("ID: %d | %s (%s)\n", p.ID, p.Nickname, p.Species)

	default:
		fmt.Printf("ID: %d | %s\n", p.ID, p.Species)
	}
}
