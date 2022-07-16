package main

import (
	"github.com/backsoul/payment/internal/api/routes"
	"github.com/backsoul/payment/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stripe/stripe-go"
)

func main() {
	stripe.Key = services.Get("STRIPE_KEY")
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers",
		AllowMethods:     "POST,PATCH",
	}))
	app.Listen(":3000")
}
