package main

import (
	"database/sql"
)

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func (c *Category) getCategoryWithName(db *sql.DB) error {
	return db.QueryRow("SELECT id, name FROM categories WHERE name=?", c.Name).Scan(&c.ID, &c.Name)
}

func (c *Category) createCategory(db *sql.DB) error {

	_, err := db.Exec("INSERT INTO categories(name) VALUES(?)", c.Name)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}
