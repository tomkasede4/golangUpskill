package models

import (
	"fmt"
)

type Book struct {
	title       string
	author      string
	publishYear int
	edition     int
	publisher   string
}

func printBooks() {
	var book1 Book
	book1.title = "Harry Potter"
	book1.author = "JK Rowling"
	book1.publishYear = 1998
	book1.edition = 1
	book1.publisher = "Penguin Publishing"

	fmt.Println("Title: ", book1.title)
	fmt.Println("Author: ", book1.author)
	fmt.Println("Publish Year: ", book1.publishYear)
	fmt.Println("Edition: ", book1.edition)
	fmt.Println("Publisher: ", book1.publisher)

}
