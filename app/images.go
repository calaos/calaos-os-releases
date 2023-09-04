package app

import (
	"github.com/calaos/calaos-os-releases/models"
	"github.com/gofiber/fiber/v2"
)

func (a *AppServer) apiV4Images(c *fiber.Ctx) (err error) {
	imgs, err := models.GetAllImages()

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		logging.Warn(err)
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"images": imgs,
	})
}

func (a *AppServer) apiV4ImagesDev(c *fiber.Ctx) (err error) {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "ok",
	})
}
