package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {

	a = App{}
	a.Initialize("root", "123!@#QWEasd", "article_maker_db")
	createTestDatabase()
	initializeTestDb()
	code := m.Run()
	clearTable()
	dropTestDb()
	os.Exit(code)
}

func TestCreateArticle(t *testing.T) {
	payload := []byte(`{
	"title": "licky licky",
	"body": "test_body",
	"category": "test_category",
	"publisher": "test_publisher",
	"published_at": "2020-02-23 22:46:00"
}`)
	req, _ := http.NewRequest("POST", "/article", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}
	if m["title"] != "licky licky" {
		t.Errorf("Expected title name to be 'licky licky'. Got '%v'", m["title"])
	}
	if m["body"] != "test_body" {
		t.Errorf("Expected body age to be 'test_body'. Got '%v'", m["body"])
	}

	if m["category"] != "test_category" {
		t.Errorf("Expected category to be 'test_category'. Got '%v'", m["category"])
	}

	if m["publisher"] != "test_publisher" {
		t.Errorf("Expected publisher to be 'test_publisher'. Got '%v'", m["publisher"])
	}

	if m["published_at"] != "2020-02-23 22:46:00" {
		t.Errorf("Expected published at to be '2020-02-23 22:46:00'. Got '%v'", m["published_at"])
	}
}

func TestGetArticle(t *testing.T) {
	payload := []byte(`{}`)
	req, _ := http.NewRequest("GET", "/article/1", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}
	if m["title"] != "licky licky" {
		t.Errorf("Expected title name to be 'licky Tad'. Got '%v'", m["title"])
	}
	if m["body"] != "test_body" {
		t.Errorf("Expected body age to be 'test_body'. Got '%v'", m["body"])
	}
	if m["category"] != "test_category" {
		t.Errorf("Expected category to be 'test_category'. Got '%v'", m["category"])
	}

	if m["publisher"] != "test_publisher" {
		t.Errorf("Expected publisher to be 'test_publisher'. Got '%v'", m["publisher"])
	}

	if m["published_at"] != "2020-02-23 22:46:00" {
		t.Errorf("Expected published at to be '2020-02-23 22:46:00'. Got '%v'", m["published_at"])
	}
}

func TestGetAllArticles(t *testing.T) {
	payload := []byte(`{}`)
	req, _ := http.NewRequest("GET", "/article", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m []map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}

	if len(m) != 1 {
		t.Errorf("Expected response size '1'. Got '%v'", len(m))
	}

}

func TestGetAllArticlesByCategory(t *testing.T) {
	payload := []byte(`{}`)
	req, _ := http.NewRequest("GET", "/article?category=test_category", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m []map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}

	if len(m) != 1 {
		t.Errorf("Expected response size '1'. Got '%v'", len(m))
	}

}

func TestGetAllArticlesByPublisher(t *testing.T) {
	payload := []byte(`{}`)
	req, _ := http.NewRequest("GET", "/article?publisher=test_publisher", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m []map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}

	if len(m) != 1 {
		t.Errorf("Expected response size '1'. Got '%v'", len(m))
	}

}

func TestGetAllArticlesByCreatedAt(t *testing.T) {
	payload := []byte(`{}`)
	req, _ := http.NewRequest("GET", "/article?published_at=2020-02-23 22:46:00", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m []map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}

	if len(m) != 1 {
		t.Errorf("Expected response size '1'. Got '%v'", len(m))
	}

}

func TestGetAllArticlesByPublishedAt(t *testing.T) {
	payload := []byte(`{}`)
	req, _ := http.NewRequest("GET", "/article?published_at=2020-02-23 22:46:00", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m []map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}

	if len(m) != 1 {
		t.Errorf("Expected response size '1'. Got '%v'", len(m))
	}

}

func TestUpdateArticle(t *testing.T) {
	payload := []byte(`{
	"id": 1,
	"title": "licky Boby",
	"body": "test_body",
	"category": "test_category",
	"publisher": "test_publisher",
	"published_at": "2020-02-23 22:46:00"
}`)
	req, _ := http.NewRequest("PUT", "/article", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}
	if m["title"] != "licky Boby" {
		t.Errorf("Expected article title to be 'licky Boby'. Got '%v'", m["title"])
	}
	if m["body"] != "test_body" {
		t.Errorf("Expected article body to be 'test_body'. Got '%v'", m["body"])
	}

	if m["category"] != "test_category" {
		t.Errorf("Expected article category to be 'test_category'. Got '%v'", m["category"])
	}

	if m["publisher"] != "test_publisher" {
		t.Errorf("Expected article publisher to be 'test_publisher'. Got '%v'", m["publisher"])
	}

	if m["published_at"] != "2020-02-23 22:46:00" {
		t.Errorf("Expected published at to be '2020-02-23 22:46:00'. Got '%v'", m["published_at"])
	}
}

func TestUpdateArticleWithNewCategory(t *testing.T) {
	payload := []byte(`{
	"id": 1,
	"title": "licky Tad",
	"body": "test_body",
	"category": "new_category",
	"publisher": "test_publisher",
	"published_at": "2020-02-23 22:46:00"
}`)
	req, _ := http.NewRequest("PUT", "/article", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}

	if m["category"] != "new_category" {
		t.Errorf("Expected article category to be 'new_category'. Got '%v'", m["category"])
	}
}

func TestUpdateArticleWithNewPublisher(t *testing.T) {
	payload := []byte(`{
	"id": 1,
	"title": "licky Tad",
	"body": "test_body",
	"category": "test_category",
	"publisher": "new_publisher",
	"published_at": "2020-02-23 22:46:00"
}`)
	req, _ := http.NewRequest("PUT", "/article", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}

	if m["publisher"] != "new_publisher" {
		t.Errorf("Expected publisher to be 'new_publisher'. Got '%v'", m["publisher"])
	}

}

func TestDeleteArticle(t *testing.T) {
	payload := []byte(`{}`)
	req, _ := http.NewRequest("DELETE", "/article/1", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err)
	}
	if m["result"] != "success" {
		t.Errorf("Expected delete response to be 'success'. Got '%v'", m["result"])
	}
}

func createTestDatabase() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	a.DB.Exec("DELETE FROM articles_test")
	a.DB.Exec("ALTER TABLE articles_test AUTO_INCREMENT = 1")
}

func dropTestDb() {
	a.DB.Exec("DROP DATABASE article_maker_test_db")
}

const tableCreationQuery = `
CREATE DATABASE IF NOT EXISTS article_maker_test_db;
`

func initializeTestDb() {

	a.Initialize("root", "123!@#QWEasd", "article_maker_test_db")

	tables := []string{
		tableArticles,
		tableCategories,
		tablePublishers,
	}

	for _, v := range tables {
		if _, err := a.DB.Exec(v); err != nil {
			log.Fatal(err)
		}
	}

}

const tableArticles = `CREATE TABLE IF NOT EXISTS articles (
  id int(11) NOT NULL AUTO_INCREMENT,
  category_id int(11) NOT NULL,
  publisher_id int(11) NOT NULL,
  title text NOT NULL,
  body text NOT NULL,
  published_at datetime DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
)`

const tableCategories = `
CREATE TABLE IF NOT EXISTS categories (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
)`

const tablePublishers = `
CREATE TABLE IF NOT EXISTS publishers (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
)
`

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
