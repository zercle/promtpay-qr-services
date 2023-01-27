package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	helpers "github.com/zercle/gofiber-helpers"
)

var ApiLimiter = limiter.New(limiter.Config{
	Max:        750,
	Expiration: 30 * time.Second,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.Get(fiber.HeaderXForwardedFor)
	},
	LimitReached: func(c *fiber.Ctx) error {
		return helpers.NewError(http.StatusTooManyRequests, helpers.WhereAmI(), http.StatusText(http.StatusTooManyRequests))
	},
})
