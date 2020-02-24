package main

import (
	"database/sql"
)

type Publisher struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func (p *Publisher) getPublisherWithName(db *sql.DB) error {
	return db.QueryRow("SELECT id, name FROM publishers WHERE name=?", p.Name).Scan(&p.ID, &p.Name)
}

func (p *Publisher) createPublisher(db *sql.DB) error {

	_, err := db.Exec("INSERT INTO publishers(name) VALUES(?)", p.Name)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}
