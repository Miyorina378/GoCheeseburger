package main

import (
	"hexagonal_cheese/handler"
	"hexagonal_cheese/packages/database"
	"hexagonal_cheese/repository"
	"hexagonal_cheese/service"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

    conn, err := database.NewMySqlConnection()
    if err != nil {
        log.Fatalf("failed to connect to MySQL (b2c): %v", err)
    }
    defer conn.Close()

    connms, err := database.NewMySqlBcloudMS()
    if err != nil {
        log.Fatalf("failed to connect to MySQL (bcloud ms): %v", err)
    }
    defer connms.Close()

	repo := repository.NewRepositories(conn, connms)
	serv := service.NewServices(repo)
	hand := handler.NewHandlers(serv)

	cheese := app.Group("/cheese")
	cheeseHandler := hand.Cheese
	cheese.Get("/get", cheeseHandler.GetAllCheese)
	cheese.Get("/get/:id", cheeseHandler.GetCheese)
	cheese.Post("/create", cheeseHandler.CreateCheese)
	cheese.Post("/update/:id", cheeseHandler.UpdateCheese)
	cheese.Delete("/delete/:id", cheeseHandler.DeleteCheese)
	cheese.Get("/getMostCountry", cheeseHandler.GetMostCountryCheese)
	app.Listen(":8080")
}
