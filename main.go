package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Inisialisasi aplikasi Fiber
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware logger
	app.Use(func(c *fiber.Ctx) error {
		log.Println("Request:", c.Method(), c.Path())
		return c.Next()
	})

	// Middleware custom
	var customMiddleware = func(c *fiber.Ctx) error {
		log.Println("Custom Middleware")
		return c.Next()
	}

	// Middleware Limiter
	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	}))

	// Middleware Helmet
	app.Use(helmet.New())

	// Middleware CORS
	app.Use(cors.New())

	// Grup routing dengan middleware
	api := app.Group("/api", customMiddleware)

	// Route GET /api/hello
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Menggunakan file statis
	app.Static("/", "./public")

	// Template engine
	app.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "My Website",
		})
	})

	// Jalankan aplikasi pada port 3000
	log.Fatal(app.Listen(":8080"))
}
