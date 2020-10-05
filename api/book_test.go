package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "CAN'T HURT ME", Author: "Goggins, David", ISBN: "978-1544512273"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"CAN'T HURT ME","author":"Goggins, David","isbn":"978-1544512273"}`, string(json), "Book Marshalling is incorrect!")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"CAN'T HURT ME","author":"Goggins, David","isbn":"978-1544512273"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "CAN'T HURT ME", Author: "Goggins, David", ISBN: "978-1544512273"}, book, "Unmarshalling is incorrect")
}
