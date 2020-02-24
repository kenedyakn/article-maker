package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) InitializeTestDB(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) handleGetArticles(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	publisher := r.URL.Query().Get("publisher")
	publishedAt := r.URL.Query().Get("published_at")
	createdAt := r.URL.Query().Get("created_at")
	articles, err := getAllArticles(a.DB, category, publisher, publishedAt, createdAt)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, articles)
}

func (a *App) handleGetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}
	article := Article{ID: id}
	if err := article.getArticle(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Article not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, article)
}

func (a *App) handleCreateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&article); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	var publisher Publisher
	publisher.Name = article.PublisherName
	if err := publisher.getPublisherWithName(a.DB); err != nil {
		if err := publisher.createPublisher(a.DB); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var category Category
	category.Name = article.CategoryName
	if err := category.getCategoryWithName(a.DB); err != nil {
		if err := category.createCategory(a.DB); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	createArticle := CreateArticle{
		Title:       article.Title,
		Body:        article.Body,
		CategoryId:  category.ID,
		PublisherId: publisher.ID,
		PublishedAt: article.PublishedAt,
	}

	if err := createArticle.createArticle(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, article)
}

func (a *App) handleUpdateArticle(w http.ResponseWriter, r *http.Request) {

	var article Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&article); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	var publisher Publisher
	publisher.Name = article.PublisherName
	if err := publisher.getPublisherWithName(a.DB); err != nil {
		if err := publisher.createPublisher(a.DB); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var category Category
	category.Name = article.CategoryName
	if err := category.getCategoryWithName(a.DB); err != nil {
		if err := category.createCategory(a.DB); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	updateArticle := UpdateArticle{
		ID:          article.ID,
		Title:       article.Title,
		Body:        article.Body,
		CategoryId:  category.ID,
		PublisherId: publisher.ID,
		PublishedAt: article.PublishedAt,
	}

	if err := updateArticle.updateArticle(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, article)
}

func (a *App) handleDeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	article := Article{ID: id}
	if err := article.deleteArticle(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := category.createCategory(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, category)
}

func (a *App) handleCreatePublisher(w http.ResponseWriter, r *http.Request) {
	var publisher Publisher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&publisher); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := publisher.createPublisher(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, publisher)
}

func (a *App) initializeRoutes() {
	// Articles
	a.Router.HandleFunc("/article", a.handleGetArticles).Methods("GET")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.handleGetArticle).Methods("GET")
	a.Router.HandleFunc("/article", a.handleCreateArticle).Methods("POST")
	a.Router.HandleFunc("/article", a.handleUpdateArticle).Methods("PUT")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.handleDeleteArticle).Methods("DELETE")

	// Categories
	a.Router.HandleFunc("/category", a.handleCreateCategory).Methods("POST")

	// Publishers
	a.Router.HandleFunc("/publisher", a.handleCreatePublisher).Methods("POST")
	// Publishers
	a.Router.HandleFunc("/publisher", a.pubName).Methods("GET")
}

func (a *App) pubName(w http.ResponseWriter, r *http.Request) {
	var publisher Publisher
	name := r.URL.Query().Get("name")

	publisher.Name = name

	if err := publisher.getPublisherWithName(a.DB); err != nil {
		if err := publisher.createPublisher(a.DB); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondWithJSON(w, http.StatusCreated, publisher)
}
