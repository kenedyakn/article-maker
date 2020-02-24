package main

import (
	"database/sql"
	"fmt"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func (a *Article) getCategory(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT title, body FROM articles WHERE id=%d", a.ID)
	return db.QueryRow(statement).Scan(&a.Title, &a.Body)
}

func (c *Category) getCategoryWithName(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, name FROM categories WHERE name='%s'", c.Name)
	return db.QueryRow(statement).Scan(&c.ID, &c.Name)
}

func (a *Article) updateCategory(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE articles SET title='%s', body='%s' WHERE id=%d", a.Title, a.Body, a.ID)
	_, err := db.Exec(statement)
	return err
}

func (a *Article) deleteCategory(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM categories WHERE id=%d", a.ID)
	_, err := db.Exec(statement)
	return err
}

func (c *Category) getCategoryByName(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id FROM categories WHERE name='%d' LIMIT 1", c.ID)
	return db.QueryRow(statement).Scan(&c.ID)
}

func (c *Category) createCategory(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO categories(name) VALUES('%s')", c.Name)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}


/*func getAllCategories(db *sql.DB, start, count int) ([]Article, error) {
	statement := fmt.Sprintf("SELECT id, category_id, publisher_id,title, body FROM articles LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	articles := []Article{}
	for rows.Next() {
		var a Article
		if err := rows.Scan(&a.ID, &a.CategoryId, &a.PublisherId, &a.Title, &a.Body); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	return articles, nil
}
*/