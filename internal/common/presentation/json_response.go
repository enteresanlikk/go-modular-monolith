package commonPresentation

import (
	"github.com/gofiber/fiber/v2"
)

func SetJsonResponseHeader(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")
}

func SetResponseStatus(ctx *fiber.Ctx, status int) {
	ctx.Status(status)
}

func SetJsonResponseBody(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(data)
}

func JsonResponse(ctx *fiber.Ctx, data interface{}) error {
	SetJsonResponseHeader(ctx)
	return SetJsonResponseBody(ctx, data)
}

func JsonResponseWithStatus(ctx *fiber.Ctx, status int, data interface{}) error {
	SetJsonResponseHeader(ctx)
	SetResponseStatus(ctx, status)
	return SetJsonResponseBody(ctx, data)
}
