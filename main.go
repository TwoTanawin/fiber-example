package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello GOJO")
	// })

	app.Get("/books", getBooks)

	app.Get("/books/:id", getBook)

	app.Listen(":8081")
}

// GET all
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

// GET one
func getBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookID {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not found eiei")
}
