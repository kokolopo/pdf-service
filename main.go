package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/set-pdf/:penetapan_id", func(c *fiber.Ctx) error {
		penetapanID := c.Params("penetapan_id")

		c.Cookie(&fiber.Cookie{
			Name:     "WhatEver",
			Value:    "0f78c540a08c25d497e42b1c52100c1875273d4906a3d172a9554c20a4da594930e3824b",
			Domain:   "simasda.kotabogor.go.id",
			Path:     "/",
			SameSite: "Lax",
		})

		// Set cookie
		c.Cookie(&fiber.Cookie{
			Name:     "auth_tkt",
			Value:    "55e243b09d393c81709a58779a4051067d5f265e0383dc194c160853eb5f2d572aaf37be3f38e2efe259af10c20fae94d70549fe84bbcda1be6f65cff6bc683b66c3463c197!userid_type:int",
			Domain:   "simasda.kotabogor.go.id",
			Path:     "/",
			SameSite: "Lax",
		})

		// Redirect to the desired URL
		// return c.Redirect("https://simasda.kotabogor.go.id/devel/aset/inventaris/kiba/lki/rpt?id=" + penetapanID)
		// return c.Redirect("https://simasda.kotabogor.go.id/inventaris/aset/inventaris/kibb/lki/rpt?id" + penetapanID)
		// https://simasda.kotabogor.go.id/inventaris/aset/inventaris/kibb/lki/rpt?id=5730827
		return c.SendString("ok " + penetapanID)
	})

	log.Fatal(app.Listen(":5001"))
}
