package controllers

import (
	"github.com/SenderAPI/mammon/pkg/types"
	"github.com/SenderAPI/mammon/pkg/utils"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/paymentmethod"
)

func CreateIdPayment(c *fiber.Ctx) error {
	id, err := gonanoid.New()
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   id,
	})
}

func CreateIntentPayment(ctx *fiber.Ctx) error {
	var payload types.Payload

	if err := ctx.BodyParser(&payload); err != nil {
		ctx.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	paymentMethodParams := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String(payload.Token),
		},
		Type: stripe.String("card"),
	}
	pyMethod, err := paymentmethod.New(paymentMethodParams)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	paymentIntentParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1000000),
		Currency: stripe.String(string(stripe.CurrencyCOP)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		StatementDescriptor: stripe.String("Custom descriptor"),
		PaymentMethod:       &pyMethod.ID,
	}
	pyIntent, err := paymentintent.New(paymentIntentParams)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   string(utils.ParseResponseJson(pyIntent)),
	})
}

func ConfirmPaymentIntent(c *fiber.Ctx) error {
	id := c.Params("id")
	params := &stripe.PaymentIntentConfirmParams{}
	piConfirm, err := paymentintent.Confirm(
		id,
		params,
	)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   piConfirm,
	})
}
