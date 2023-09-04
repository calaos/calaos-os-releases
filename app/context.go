package app

import (
	"strings"

	"github.com/calaos/calaos-os-releases/config"
	"github.com/gofiber/fiber/v2"
)

const (
	TOKEN_FILE = "/run/calaos/calaos-ct.token"
)

var (
	tokenStr string
)

func NewTokenMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token string

		//get token
		headerValue := c.Get("authorization")
		if headerValue != "" {
			split := strings.SplitN(headerValue, " ", 2)

			if len(split) == 2 && strings.ToLower(split[0]) == "bearer" {
				token = split[1]
			}
		}

		tokenStr = strings.TrimSpace(config.Config.String("general.update-token"))

		if tokenStr == "" {
			logging.Fatal("update-token is not set in config!")
		}

		//Check token validity
		if token != tokenStr {
			c.Locals("token-valid", false)
		} else {
			c.Locals("token-valid", true)
		}

		return c.Next()
	}
}
