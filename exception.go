package xfiber

import "github.com/gofiber/fiber/v3"

const (
	StatusAuthorizedExpired = 406
)

// HTTP异常
func HTTPException(c fiber.Ctx, statusCode int, detail string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"detail": detail,
	})
}

// 认证失败
func AuthenticationError(c fiber.Ctx) error {
	return HTTPException(c, fiber.StatusUnauthorized, "认证失败")
}

// 认证过期
func AuthenticationExpired(c fiber.Ctx) error {
	return HTTPException(c, StatusAuthorizedExpired, "认证过期")
}

// 错误请求
func HTTPExceptionBadRequest(c fiber.Ctx, message string) error {
	return HTTPException(c, fiber.StatusBadRequest, message)
}

// 禁止访问
func HTTPExceptionForbidden(c fiber.Ctx) error {
	return HTTPException(c, fiber.StatusForbidden, "禁止访问")
}

// 请求参数错误
func HTTPExceptionUnprocessableEntity(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"detail": "请求参数错误(" + err.Error() + ")",
		"data":   err.Error(),
	})
}

// 数据库错误
func HTTPExceptionDatabase(c fiber.Ctx) error {
	return HTTPException(c, fiber.StatusServiceUnavailable, "数据库错误")
}

// 内部错误
func HTTPExceptionInternalServerError(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"detail": "内部错误",
		"data":   err.Error(),
		"stack":  c.Locals("stack"),
	})
}
