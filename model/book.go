package model

type Book struct {
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}

type Books []Book
