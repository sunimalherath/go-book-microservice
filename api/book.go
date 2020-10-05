package api

import (
	"encoding/json"
	"net/http"
)

// Book struct with Author, Book name and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// ToJSON for marshalling book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return ToJSON
}

// FromJSON() for unmarshalling book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}

	return book
}

// Sample Books slice to test the API
var Books = []Book{
	Book{Title: "CAN'T HURT ME", Author: "Goggins, David", ISBN: "978-1544512273"},
	Book{Title: "Atomic Habits", Author: "Clear, James", ISBN: "978-1847941831"},
}

// BookHandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}
