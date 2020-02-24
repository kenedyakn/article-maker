package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type Article struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	CategoryName  string `json:"category"`
	PublisherName string `json:"publisher"`
	PublishedAt   string `json:"published_at"`
	CreatedAt     string `json:"created_at"`
}

type UpdateArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	CategoryId  int    `json:"category_id"`
	PublisherId int    `json:"publisher_id"`
	PublishedAt string `json:"published_at"`
}

type CreateArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	CategoryId  int    `json:"category_id"`
	PublisherId int    `json:"publisher_id"`
	PublishedAt string `json:"published_at"`
}

func (a *CreateArticle) createArticle(db *sql.DB) error {

	_, err := db.Exec("INSERT INTO articles(category_id, publisher_id,title,body,published_at) VALUES(?,?,?,?,?)",
		a.CategoryId, a.PublisherId, a.Title, a.Body, a.PublishedAt)

	if err != nil {
		return err
	}

	return nil
}

func (a *UpdateArticle) updateArticle(db *sql.DB) error {
	statement := fmt.Sprintf(`
	UPDATE articles SET title='%s', body='%s', category_id=%d, publisher_id=%d, published_at='%s' WHERE id=%d`,
		a.Title, a.Body, a.CategoryId, a.PublisherId, a.PublishedAt, a.ID)
	_, err := db.Exec(statement)
	return err
}

func (a *Article) getArticle(db *sql.DB) error {

	statement := fmt.Sprintf(`
    SELECT a.id, a.title, a.body,c.name, p.name, a.published_at, a.created_at
    FROM articles a
    LEFT JOIN categories c ON a.category_id = c.id 
	LEFT JOIN publishers p ON a.publisher_id = p.id
    WHERE a.id = %d
  `, a.ID)
	return db.QueryRow(statement).Scan(&a.ID, &a.Title, &a.Body, &a.CategoryName, &a.PublisherName, &a.PublishedAt, &a.CreatedAt)
}

func (a *Article) updateArticle(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE articles SET title='%s', body='%s' WHERE id=%d", a.Title, a.Body, a.ID)
	_, err := db.Exec(statement)
	return err
}

func (a *Article) deleteArticle(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM articles WHERE id=%d", a.ID)
	_, err := db.Exec(statement)
	return err
}

func getAllArticles(db *sql.DB, category string, publisher string, publishedAt string, createdAt string) ([]Article, error) {
	statement := fmt.Sprintf(`
    SELECT a.id, a.title, a.body,c.name, p.name, a.published_at, a.created_at
    FROM articles a
    LEFT JOIN categories c ON a.category_id = c.id 
	LEFT JOIN publishers p ON a.publisher_id = p.id `)

	// Add query conditions if any of the query parameters is received
	if category != "" || publisher != "" || publishedAt != "" || createdAt != "" {

		statement += "WHERE "

		values := []string{}
		if category != "" {
			values = append(values, fmt.Sprintf(`c.name = '%s'`, category))
		}

		if publisher != "" {
			values = append(values, fmt.Sprintf(`p.name = '%s'`, publisher))
		}

		if publishedAt != "" {
			values = append(values, fmt.Sprintf(`a.published_at = '%s'`, publishedAt))
		}

		if createdAt != "" {
			values = append(values, fmt.Sprintf(`a.created_at = '%s'`, createdAt))
		}

		result := strings.Join(values, " OR ")

		statement += result
	}

	rows, err := db.Query(statement)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	articles := []Article{}
	for rows.Next() {
		var a Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Body, &a.CategoryName, &a.PublisherName, &a.PublishedAt, &a.CreatedAt); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	return articles, nil
}
