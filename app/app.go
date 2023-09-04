package app

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/calaos/calaos-os-releases/config"
	logger "github.com/calaos/calaos-os-releases/log"
	"github.com/sirupsen/logrus"
)

const (
	maxFileSize = 1 * 1024 * 1024 * 1024
)

type AppServer struct {
	quitHeartbeat chan interface{}
	wgDone        sync.WaitGroup

	appFiber *fiber.App
}

var logging *logrus.Entry

func init() {
	logging = logger.NewLogger("app")
}

// Init the app
func NewApp() (a *AppServer, err error) {
	logging.Infoln("Init server")

	a = &AppServer{
		quitHeartbeat: make(chan interface{}),
		appFiber: fiber.New(fiber.Config{
			ServerHeader:          "Calaos OS Releases (Linux)",
			ReadTimeout:           time.Second * 20,
			AppName:               "Calaos OS Releases",
			DisableStartupMessage: true,
			EnablePrintRoutes:     false,
			BodyLimit:             maxFileSize,
		}),
	}

	a.appFiber.
		Use(fiberLog.New(fiberLog.Config{})).
		Use(NewTokenMiddleware())

	a.appFiber.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1",
	}))

	a.appFiber.Hooks().OnShutdown(func() error {
		a.wgDone.Done()
		return nil
	})

	a.appFiber.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://calaos.fr")
	})

	//API
	api := a.appFiber.Group("/v4")

	api.Get("/images", func(c *fiber.Ctx) error {
		return a.apiV4Images(c)
	})

	api.Get("/images-dev", func(c *fiber.Ctx) error {
		return a.apiV4ImagesDev(c)
	})

	api.Post("/image/:name", func(c *fiber.Ctx) error {
		return a.apiV4ImageSet(c)
	})

	api.Post("/image-dev/:name", func(c *fiber.Ctx) error {
		return a.apiV4ImageDevSet(c)
	})

	return
}

// Run the app
func (a *AppServer) Start() {
	addr := config.Config.String("general.address") + ":" + strconv.Itoa(config.Config.Int("general.port"))

	logging.Infoln("\u21D2 Server listening on", addr)

	go func() {
		if err := a.appFiber.Listen(addr); err != nil {
			logging.Fatalf("Failed to listen http server: %v", err)
		}
	}()
	a.wgDone.Add(1)
}

// Stop the app
func (a *AppServer) Shutdown() {
	close(a.quitHeartbeat)
	a.appFiber.Shutdown()
	a.wgDone.Wait()
}
