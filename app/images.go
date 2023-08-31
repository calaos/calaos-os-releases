package app

import (
	"github.com/gofiber/fiber/v2"
)

func (a *AppServer) apiV4Images(c *fiber.Ctx) (err error) {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "ok",
	})
}

func (a *AppServer) apiV4ImagesDev(c *fiber.Ctx) (err error) {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "ok",
	})
}
