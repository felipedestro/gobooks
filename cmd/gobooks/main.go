package main

import (
	"database/sql"
	"gobooks/internal/services"
	"gobooks/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := services.NewBookService(db)
	bookHandlers := web.NewBookHandler(bookService)

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookByID)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}
