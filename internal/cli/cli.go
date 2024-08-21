package cli

import (
	"fmt"
	"gobooks/internal/services"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	service *services.BookService
}

func NewBookCLI(service *services.BookService) *BookCLI {
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usagem: books <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simate <book_id> <book_id> <book_id> ...")
			return
		}
		booksID := os.Args[2:]
		cli.SimulateReading(booksID)
	}
}

func (cli *BookCLI) searchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)
	if err != nil {
		fmt.Println("Error searching books: ", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("no books found")
		return
	}

	fmt.Printf("%d books found\n", len(books))
	for _, book := range books {
		fmt.Printf("ID %d, Title: %s, Author: %s, Genre: %s\n",
			book.ID, book.Title, book.Author, book.Genre,
		)
	}
}

func (cli *BookCLI) SimulateReading(booksIDsStr []string) {
	var booksIDs []int
	for _, idStr := range booksIDsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid book ID: ", idStr)
			continue
		}
		booksIDs = append(booksIDs, id)
	}
	responses := cli.service.SimulateMultileReading(booksIDs, 2*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}
