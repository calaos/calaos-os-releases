package app

import (
	"github.com/calaos/calaos-os-releases/models"
	"github.com/gofiber/fiber/v2"
)

func (a *AppServer) apiV4Images(c *fiber.Ctx) (err error) {
	imgs, err := models.GetAllImages(models.Release)

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
	imgs, err := models.GetAllImages(models.Development)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		logging.Warn(err)
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"images": imgs,
	})
}

func (a *AppServer) apiV4ImageSet(c *fiber.Ctx) (err error) {
	return a.imageSet(c, models.Release)
}

func (a *AppServer) apiV4ImageDevSet(c *fiber.Ctx) (err error) {
	return a.imageSet(c, models.Development)
}

func (a *AppServer) imageSet(c *fiber.Ctx, repoType string) (err error) {
	tokenValid := c.Locals("token-valid")
	tv, ok := tokenValid.(bool)
	if !ok || !tv {
		return &fiber.Error{
			Code:    fiber.ErrUnauthorized.Code,
			Message: "token not valid",
		}
	}

	name := c.Params("name", "")
	if name == "" {
		c.Status(fiber.StatusBadRequest)
		logging.Warn("name parameter empty")
		return nil
	}

	img := new(models.Image)
	if err := c.BodyParser(img); err != nil {
		logging.Warnf("bad Image JSON: %v", err)
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	img.RepoType = repoType
	err = models.StoreImage(img)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	c.Status(fiber.StatusOK)
	return nil
}
