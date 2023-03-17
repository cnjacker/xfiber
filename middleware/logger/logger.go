package logger

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var (
	TimeZone = "Asia/Shanghai"
)

func New() fiber.Handler {
	config := logger.ConfigDefault
	config.Format = "${pid} | ${time} | ${ip}:${port} | ${status} - ${latency} ${method} ${path}\n"
	config.TimeFormat = time.DateTime
	config.TimeZone = TimeZone

	return logger.New(config)
}
