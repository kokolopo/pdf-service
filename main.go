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
			Value: "62fd5c685d450a356af16ccd56392496be85e35c5c5f5d1ceeede74bd50992b931a269a940742646f1feea8bb902cebd720c6ac50d78b469e6d8f7abf59ac96166bc51b7197!userid_type:int",
			// Value:  "8dde60729f8adb12881f9de0675d5325628b6762d25a967416f2fa7b980368f499a47eda6897fe205ae7bf41e07f889038a3af622de83cd1259da0f7b0b08b9f66bbae76197!userid_type:int",
			Domain: "simasda.kotabogor.go.id",
			Path:   "/",
		})

		// Redirect to the desired URL
		return c.Redirect("https://simasda.kotabogor.go.id/devel/aset/inventaris/kiba/lki/rpt?id=" + penetapanID)
	})

	log.Fatal(app.Listen(":6000"))
}
