package db

import "database/sql"

var migration001 = Migration{
	version: 1,
	up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`CREATE TABLE growing_areas (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			type TEXT NOT NULL,
			location_notes TEXT,
			created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now'))
		)`)

		if err != nil {
			return err
		}

		_, err = tx.Exec(`CREATE TABLE plants (
			id INTEGER PRIMARY KEY,
			species TEXT NOT NULL,
			nickname TEXT,
			variety TEXT,
			created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now'))
		)`)

		if err != nil {
			return err
		}

		_, err = tx.Exec(`CREATE TABLE plantings (
			id INTEGER PRIMARY KEY,
			plant_id INTEGER NOT NULL,
			growing_area_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL DEFAULT 1,
			planted_at TEXT NOT NULL,
			expected_harvest_at TEXT NOT NULL,
			status TEXT CHECK(status IN ('active', 'harvested', 'failed')),
			created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			FOREIGN KEY (plant_id) REFERENCES plants(id),
			FOREIGN KEY (growing_area_id) REFERENCES growing_areas(id)
		)`)

		if err != nil {
			return err
		}

		_, err = tx.Exec(`CREATE TABLE care_events (
			id INTEGER PRIMARY KEY,
			planting_id INTEGER,
			growing_area_id INTEGER,
			event_type TEXT CHECK(event_type IN ('water', 'fertilize', 'prune', 'harvest', 'weed', 'other')),
			notes TEXT,
			occurred_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			FOREIGN KEY (planting_id) REFERENCES plantings(id),
			FOREIGN KEY (growing_area_id) REFERENCES growing_areas(id),
			CHECK (
				(planting_id IS NOT NULL AND growing_area_id IS NULL) OR
				(planting_id IS NULL AND growing_area_id IS NOT NULL)
			)
		)`)

		if err != nil {
			return err
		}

		_, err = tx.Exec(`CREATE TABLE care_schedules (
			id INTEGER PRIMARY KEY,
			planting_id INTEGER,
			growing_area_id INTEGER,
			event_type TEXT CHECK(event_type IN ('water', 'fertilize', 'prune', 'harvest', 'weed', 'other')),
			frequency TEXT CHECK(frequency IN ('daily', 'weekly', 'monthly', 'custom')),
			start_date TEXT NOT NULL,
			end_date TEXT NOT NULL,
			created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
			FOREIGN KEY (planting_id) REFERENCES plantings(id) ON DELETE CASCADE,
			FOREIGN KEY (growing_area_id) REFERENCES growing_areas(id) ON DELETE CASCADE,
			CHECK (
				(planting_id IS NOT NULL AND growing_area_id IS NULL) OR
				(planting_id IS NULL AND growing_area_id IS NOT NULL)
			)
		)`)

		if err != nil {
			return err
		}

		return nil
	},
}
