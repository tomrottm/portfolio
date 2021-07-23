package book

import (
	"fiber-gorm-books/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
	Desc   string `json:"desc"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []Book
	result := db.Find(&books)
	if result.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No records found",
		})
	}

	return c.JSON(books)
}

//demo
func DisplayAllBooks(c *fiber.Ctx) error {

	var books []Book
	db := database.DB

	err := db.Find(&books).Scan(&books).Error
	if err != nil {
		return err
	}

	data := books

	return c.Render("book/index", fiber.Map{
		"Title": "All Books",
		"data":  data,
	})
}

func GetBook(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var book Book
	err := db.Find(&book, id).Scan(&book).Error
	if err != nil {
		return err
	}

	data := book

	return c.Render("book/book", fiber.Map{
		"data": data,
	})
}
func NewBook(c *fiber.Ctx) error {
	db := database.DB

	// For Manual Input
	// var book Book
	// book.Title = "Harry Potter and the Sorcerer's Stone (Harry Potter, #1)"
	// book.Author = "J.K. Rowling"
	// book.Rating = 7

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&book)

	return c.JSON(book)
}
func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("Update Book.")
}
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var book Book
	result := db.First(&book, id)

	if result.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Record not found",
		})
	}

	db.Delete(&book)
	c.Status(fiber.StatusNoContent)

	return nil
}
