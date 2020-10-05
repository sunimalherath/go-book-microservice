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
	return Book{}
}

// BookHandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {

}
