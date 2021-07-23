package main

import (
	"fmt"

	"fiber-gorm-books/book"
	"fiber-gorm-books/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	initDatabase()

	app.Get("/", helloWorld)
	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	booksRoute := app.Group("/books")
	booksRoute.Get("/", book.GetBooks)
	booksRoute.Post("/", book.NewBook)
	booksRoute.Get("/:id", book.GetBook)
	booksRoute.Delete("/:id", book.DeleteBook)
	booksRoute.Patch("/:id", book.UpdateBook)

	//demo
	app.Get("/allbooks", book.DisplayAllBooks)

	// try: http://127.0.0.1:3000/allbooks
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World! 👋")
}

func initDatabase() {
	if err := database.Open(); err != nil {
		fmt.Println("Could not open Database Connection.")
	}

	database.DB.AutoMigrate(&book.Book{})
	fmt.Println("Database successfully auto-migrated.")
}
