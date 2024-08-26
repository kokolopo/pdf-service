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
			Value: "0c599c6ca8cdfca8ae605f0a32cca214e57c1473fad351a6a1627a5c4c457a3c3c7db4bea782ddefc43b041af026c63ef8a1571c32caeca78a363563841d875066cbf95f3!userid_type:int",
			// Domain:   "simasda.kotabogor.go.id",
			Path:     "/",
			SameSite: "Lax",
			HTTPOnly: true,
		})

		// Redirect to the desired URL
		// return c.Redirect("https://simasda.kotabogor.go.id/devel/aset/inventaris/kiba/lki/rpt?id=" + penetapanID)
		return c.Redirect("https://simasda.kotabogor.go.id/inventaris/aset/inventaris/kibb/lki/rpt?id=" + penetapanID)
		// https://simasda.kotabogor.go.id/inventaris/aset/inventaris/kibb/lki/rpt?id=5730827
		// return c.SendString("ok " + penetapanID)
	})

	log.Fatal(app.Listen(":5001"))
}
