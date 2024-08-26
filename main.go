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
			Value: "205e856cb5a1801ceef7cb83f4f5b49de153019f52f5fa3c0f07f369bf6f6bec92a52e90e29028a8b3c89ba27c2b5375516e25a5b67eae36d230d5eee870bc9066cbdbf3197!userid_type:int",
			// Domain:   "simasda.kotabogor.go.id",
			Path:     "/",
			SameSite: "Lax",
			HTTPOnly: true,
		})

		// Redirect to the desired URL
		return c.Redirect("https://simasda.kotabogor.go.id/devel/aset/inventaris/kiba/lki/rpt?id=" + penetapanID)
		// return c.Redirect("https://simasda.kotabogor.go.id/inventaris/aset/inventaris/kibb/lki/rpt?id" + penetapanID)
		// https://simasda.kotabogor.go.id/inventaris/aset/inventaris/kibb/lki/rpt?id=5730827
		// return c.SendString("ok " + penetapanID)
	})

	log.Fatal(app.Listen(":5001"))
}
