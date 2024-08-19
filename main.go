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
			Value: "a2082928acb5ab25d7200c00bb534c9458c25b3621763989bbaef712dd80d806081773d815436bf89ccbf6220b1e67befb5788fcab4db5d5a007a7a7e40274dc66c32a39197!userid_type:int",
			// Value:  "8dde60729f8adb12881f9de0675d5325628b6762d25a967416f2fa7b980368f499a47eda6897fe205ae7bf41e07f889038a3af622de83cd1259da0f7b0b08b9f66bbae76197!userid_type:int",
			Domain: "simasda.kotabogor.go.id",
			Path:   "/",
		})

		// Redirect to the desired URL
		return c.Redirect("https://simasda.kotabogor.go.id/devel/aset/inventaris/kiba/lki/rpt?id=" + penetapanID)
	})

	log.Fatal(app.Listen(":5001"))
}
