package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/saiful-4321/crud/pkg/models"
)

var newBook models.Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	response, _ := json.Marshal((newBooks))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
