package main

import (
    "fmt"
)

type Book struct {
    ID int
    Title string
    Author string
    Description string
    Price float64
}

func getBookDetails(id int) *Book {
    //simulate fetching a book from db
    return &Book {
        ID: id,
        Title: "Sample Book",
        Author: "John Doe",
        Description: "Sit cupiditate ad voluptates cumque quo? Modi earum dolorem dolorem dolore itaque temporibus.",
        Price: 20.11,
    }
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        genre := r.URL.Query().Get("genre")
        if genre != "" {
            fmt.Fprintf(w, "Showing books in genre: %s", genre)
        } else {
            fmt.Fprintf(w, "Showing all books")
        }
    case "POST":
        fmt.Fprintf(w, "Adding a book")
    default:
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
    }
}

// or belowu code can be used for handling different methods

func getBookHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Getting book details")
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Adding a book")
}

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Updating book details")
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Deleting a book")
}

/* the HEAD method is similar to GET except that the server will not return a message body in the response. It is used for obtaining metadata */
func headBookHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Length", "0")
}

// OPTIONS method is used to describe the communication options for the target resource
func optionsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Allow", "GET, POST, HEAD")
    fmt.Fprintf(w, "Provides options for the resource")
}

func main() {
    book := getBookDetails(1)
    fmt.Println("Book ID: ", book.ID)
    fmt.Println("Book Title: ", book.Title)
    fmt.Println("Book Author: ", book.Author)
    fmt.Println("Book Description: ", book.Description)
    fmt.Println("Book Price: ", book.Price)
}
