package repository

import (
	"database/sql"
	"errors"

	"github.com/haley-marie/greenhousedashboard/backend/models"
)

type PlantRepository struct {
	db *sql.DB
}

func (r *PlantRepository) AddPlant(p *models.Plant) (*models.Plant, error) {
	if p.Species == "" {
		return nil, errors.New("must specify species")
	}

	sqlStatement := `
		INSERT INTO plants (species, nickname, variety)
		VALUES (?, ?, ?)
		`

	result, err := r.db.Exec(
		sqlStatement,
		p.Species,
		p.Nickname,
		p.Variety,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow("SELECT id, species, nickname, variety, created_at, updated_at FROM plants WHERE id = ?", id)

	err = row.Scan(
		&p.ID,
		&p.Species,
		&p.Nickname,
		&p.Variety,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("plant not found in database")
		}
		return nil, err
	}

	return p, nil
}
