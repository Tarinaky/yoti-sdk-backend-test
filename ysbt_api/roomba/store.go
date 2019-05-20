package roomba

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
)

func (this *Room) Store (db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var room_id uuid.UUID
	if err := tx.QueryRow(
		`INSERT INTO roomba.rooms (
			room_width,
			room_height
		) VALUES ($1, $2)
		RETURNING id`,
		this.Width,
		this.Height,
	).Scan(&room_id); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatal(err)
		}
		return err
	}

	for dirt := range this.DirtPatches {
		if _,err := tx.Exec(
			`INSERT INTO roomba.dirt (
				room_id,
				x_pos,
				y_pos
			) VALUES ($1,$2,$3)`,
			room_id,
			dirt.X,
			dirt.Y,
		); err != nil {
			if err := tx.Rollback(); err != nil {
				log.Fatal(err)
			}
			return err
		}
	}

	if _,err := tx.Exec(
		`INSERT INTO roomba.prior_runs (
			room_id,
			start_x,
			start_y,
			instructions,
			finish_x,
			finish_y,
			dirt_collected
		) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		room_id,
		this.Roomba.StartX,
		this.Roomba.StartY,
		this.Instructions,
		this.Roomba.CurrentX,
		this.Roomba.CurrentY,
		this.Roomba.DirtCollected,
	); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatal(err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

