package main

import (
	"os"

	//"github.com/xhadzi01/todobhv/taskManagement"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/xhadzi01/todobhv/controllers"
	"github.com/xhadzi01/todobhv/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".tmpl")

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")

	// REST Routing
	app.Get("/api/tasks", controllers.FetchTasks)
	app.Post("/api/tasks", controllers.CreateTask)
	app.Get("/api/tasks/:id", controllers.FetchTask)
	app.Delete("/api/tasks/:id", controllers.DeleteTask)

	// HTTP frontend Routing
	frontendRoutes := []string{
		"/",
		"/about",
	}

	for _, route := range frontendRoutes {
		app.Get(route, controllers.Home)
	}

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
