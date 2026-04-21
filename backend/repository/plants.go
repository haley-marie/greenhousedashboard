package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/haley-marie/greenhousedashboard/backend/models"
)

type PlantRepository struct {
	DB *sql.DB
}

var ErrNotFound = errors.New("not found")

func convertIfValidStr(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}

	return ""
}

func parseTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339Nano, s)
	return t, err
}

func (r *PlantRepository) AddPlant(p *models.Plant) (*models.Plant, error) {
	if p.Species == "" {
		return nil, errors.New("must specify species")
	}

	sqlStatement := `
		INSERT INTO plants (species, nickname, variety)
		VALUES (?, ?, ?)
		`

	result, err := r.DB.Exec(
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

	row := r.DB.QueryRow("SELECT id, species, nickname, variety, created_at, updated_at FROM plants WHERE id = ?", id)

	var nickname, variety sql.NullString
	var createdAtStr, updatedAtStr string

	err = row.Scan(
		&p.ID,
		&p.Species,
		&nickname,
		&variety,
		&createdAtStr,
		&updatedAtStr,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	createdAt, err := parseTime(createdAtStr)
	if err != nil {
		return nil, err
	}

	updatedAt, err := parseTime(updatedAtStr)
	if err != nil {
		return nil, err
	}

	p.CreatedAt = createdAt
	p.UpdatedAt = updatedAt
	p.Nickname = convertIfValidStr(nickname)
	p.Variety = convertIfValidStr(variety)

	return p, nil
}

func (r *PlantRepository) ListPlants() ([]models.Plant, error) {
	rows, err := r.DB.Query("SELECT * FROM plants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []models.Plant

	for rows.Next() {
		var p models.Plant
		var nickname, variety sql.NullString
		var createdAtStr, updatedAtStr string

		err = rows.Scan(
			&p.ID,
			&p.Species,
			&nickname,
			&variety,
			&createdAtStr,
			&updatedAtStr,
		)
		if err != nil {
			return nil, err
		}

		createdAt, err := parseTime(createdAtStr)
		if err != nil {
			return nil, err
		}

		updatedAt, err := parseTime(updatedAtStr)
		if err != nil {
			return nil, err
		}

		p.CreatedAt = createdAt
		p.UpdatedAt = updatedAt
		p.Nickname = convertIfValidStr(nickname)
		p.Variety = convertIfValidStr(variety)

		plants = append(plants, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return plants, nil
}

func (r *PlantRepository) GetPlantByID(id int) (*models.Plant, error) {
	row := r.DB.QueryRow(`
	SELECT id, species, nickname, variety, created_at, updated_at 
	FROM plants
	WHERE id = ?
	`, id)

	p := &models.Plant{}

	err := row.Scan(
		&p.ID,
		&p.Species,
		&p.Nickname,
		&p.Variety,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return p, nil
}
