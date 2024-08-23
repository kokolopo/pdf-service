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
			Value: "7f2617c8d555db55b6b43908c847e3804cac5263e7d255a5a6d3ea9c2824764f3fd846e3e3a043d18f86e3981180adf9f4ae384e993a5668ea2eb825469b874866c826bb197!userid_type:int",
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
