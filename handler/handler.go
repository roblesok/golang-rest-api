package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/roblesok/golang-rest-api/model"
)

type bookHandler struct {
	sync.Mutex
	books model.Books
}

func NewBookHandler() *bookHandler {
	return &bookHandler{
		books: model.Books{
			model.Book{"Animal Farm", "Political", "George Orwell"},
			model.Book{"Nineteen Eighty-Four", "Fiction, Social SPy", "George Orwell"},
		},
	}
}

func (bh *bookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bh.get(w, r)
	default:
		SendErr(w, http.StatusMethodNotAllowed, "Method not allowed!")
	}
}

func (bh *bookHandler) get(w http.ResponseWriter, r *http.Request) {
	defer bh.Unlock()
	bh.Lock()
	SendJSON(w, http.StatusOK, bh.books)
}

func SendErr(w http.ResponseWriter, code int, msg string) {
	SendJSON(w, code, map[string]string{"error": msg})
}

func SendJSON(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
