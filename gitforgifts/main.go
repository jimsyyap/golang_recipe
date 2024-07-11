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

func main() {
    book := getBookDetails(1)
    fmt.Println("Book ID: ", book.ID)
    fmt.Println("Book Title: ", book.Title)
    fmt.Println("Book Author: ", book.Author)
    fmt.Println("Book Description: ", book.Description)
    fmt.Println("Book Price: ", book.Price)
}
