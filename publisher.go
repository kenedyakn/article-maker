package main

import (
	"database/sql"
	"fmt"
)

type Publisher struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func (p *Publisher) getPublisher(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, name FROM publishers WHERE id=%d", p.ID)
	return db.QueryRow(statement).Scan(&p.ID, &p.Name)
}

func (p *Publisher) getPublisherWithName(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, name FROM publishers WHERE name='%s'", p.Name)
	return db.QueryRow(statement).Scan(&p.ID, &p.Name)
}

func (a *Article) updatePublisher(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE publishers SET title='%s', body='%s' WHERE id=%d", a.Title, a.Body, a.ID)
	_, err := db.Exec(statement)
	return err
}

func (a *Article) deletePublisher(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM publishers WHERE id=%d", a.ID)
	_, err := db.Exec(statement)
	return err
}

func (p *Publisher) createPublisher(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO publishers(name) VALUES('%s')", p.Name)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getAllPublishers(db *sql.DB, start, count int) ([]Publisher, error) {
	statement := fmt.Sprintf("SELECT id,name,created_at FROM publishers LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	publishers := []Publisher{}
	for rows.Next() {
		var p Publisher
		if err := rows.Scan(&p.ID, &p.Name, &p.CreatedAt); err != nil {
			return nil, err
		}
		publishers = append(publishers, p)
	}

	return publishers, nil
}
