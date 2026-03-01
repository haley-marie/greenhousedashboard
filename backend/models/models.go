package models

import "time"

type Plant struct {
	ID        int
	Species   string
	Nickname  string
	Variety   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GrowingArea struct {
	ID            int
	Name          string
	AreaType      string
	LocationNotes string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Planting struct {
	ID                int
	PlantID           int
	GrowingAreaID     int
	Quantity          int
	PlantedAt         time.Time
	ExpectedHarvestAt time.Time
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type CareEvent struct {
	ID            int
	PlantingID    int
	GrowingAreaID int
	EventType     string
	Notes         string
	OccurredAt    time.Time
	CreatedAt     time.Time
}

type CareSchedule struct {
	ID            int
	PlantingID    int
	GrowingAreaID int
	EventType     string
	Frequency     string
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
