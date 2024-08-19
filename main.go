package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/set-pdf/:penetapan_id", func(c *fiber.Ctx) error {
		penetapanID := c.Params("penetapan_id")

		// Set cookie
		c.Cookie(&fiber.Cookie{
			Name:  "auth_tkt",
			Value: "765f33ff08199a88d15e345708234871fd9131200b90e4c0031bf7055f08ebac59daace98e2fc091e6bc32308e91a03c0b8b1eb01372cb0dde1b71e563e676fe66c332ed197!userid_type:int",
			// Value:  "8dde60729f8adb12881f9de0675d5325628b6762d25a967416f2fa7b980368f499a47eda6897fe205ae7bf41e07f889038a3af622de83cd1259da0f7b0b08b9f66bbae76197!userid_type:int",
			Domain: "simasda.kotabogor.go.id",
			Path:   "/",
		})

		// Redirect to the desired URL
		return c.Redirect("https://simasda.kotabogor.go.id/devel/aset/inventaris/kiba/lki/rpt?id=" + penetapanID)
	})

	log.Fatal(app.Listen(":5001"))
}
