package xfiber

import (
	"github.com/cnjacker/xfiber/middleware/auth"
	"github.com/cnjacker/xfiber/middleware/logger"
	"github.com/cnjacker/xfiber/middleware/operation"
	"github.com/cnjacker/xfiber/middleware/recover"
	"github.com/cnjacker/xfiber/middleware/signature"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

var (
	BodyLimit = 64 * 1024 * 1024
)

func New() *fiber.App {
	config := fiber.Config{
		BodyLimit: BodyLimit,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			if err, ok := err.(*fiber.Error); ok {
				switch err.Code {
				case fiber.StatusNotFound:
					return HTTPException(c, err.Code, "资源不存在")
				case fiber.StatusInternalServerError:
					return HTTPException(c, err.Code, "内部错误")
				default:
					return HTTPExceptionInternalServerError(c, err)
				}
			}

			return HTTPExceptionInternalServerError(c, err)
		},
	}

	app := fiber.New(config)

	return app
}

func NewWithMiddleware() *fiber.App {
	app := New()

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(signature.New())
	app.Use(auth.New())
	app.Use(operation.New())

	return app
}
