package handlers

import (
	"fmt"
	"go-books/data"
	"log"
	"net/http"
)

type Env struct {
	Books data.BookModel
}

// Define GetAll as a method on Env.
func (env *Env) GetAll(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	bks, err := env.Books.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
