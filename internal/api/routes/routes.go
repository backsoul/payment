package routes

import (
	"github.com/backsoul/payment/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	orders := app.Group("/orders")
	orders.Post("/", controllers.CreateIdPayment)
	orders.Post("/:id", controllers.CreateIntentPayment)
	orders.Post("/confirm/:id", controllers.ConfirmPaymentIntent)
}
