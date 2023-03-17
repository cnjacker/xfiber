package limiter

import (
	"fmt"
	"time"

	"github.com/cnjacker/xfiber"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"

	mapset "github.com/deckarep/golang-set/v2"
)

var (
	Expiration = 2 * time.Second
)

func New() fiber.Handler {
	return limiter.New(
		limiter.Config{
			Next: func(c fiber.Ctx) bool {
				return !mapset.NewSet("POST").Contains(c.Method())
			},
			Max: 1,
			KeyGenerator: func(c fiber.Ctx) string {
				return fmt.Sprintf("%s:%s:%s", c.IP(), c.Port(), c.Path())
			},
			Expiration: Expiration,
			LimitReached: func(c fiber.Ctx) error {
				return xfiber.HTTPException(
					c, fiber.StatusTooManyRequests, "操作太过频繁, 请稍后再试",
				)
			},
			SkipFailedRequests: true,
		},
	)
}
