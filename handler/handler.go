package handler

import (
	"net/http"
	"sync"

	"github.com/roblesok/golang-rest-api/model"
	"github.com/roblesok/golang-rest-api/utils"
)

// BookHandler represents the /api/books
type BookHandler struct {
	sync.Mutex
	books model.Books
}

// NewBookHandler returns a *bookHandler
func NewBookHandler() *BookHandler {
	return &BookHandler{
		books: model.Books{
			model.Book{Title: "Animal Farm", Genre: "Political", Author: "George Orwell"},
			model.Book{Title: "Nineteen Eighty-Four", Genre: "Fiction, Social SPy", Author: "George Orwell"},
		},
	}
}

func (bh *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bh.get(w, r)
	default:
		utils.SendErr(w, http.StatusMethodNotAllowed, "Method not allowed!")
	}
}

func (bh *BookHandler) get(w http.ResponseWriter, r *http.Request) {
	defer bh.Unlock()
	bh.Lock()
	id, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.SendJSON(w, http.StatusOK, bh.books)
		return
	}
	if id >= len(bh.books) || id < 0 {
		utils.SendErr(w, http.StatusNotFound, "Not Found")
		return
	}
	utils.SendJSON(w, http.StatusOK, bh.books[id])
}
