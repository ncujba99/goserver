package rest

import "github.com/gofiber/fiber/v2"

func index_handeler(ctx *fiber.Ctx) error {
	res := map[string]any{}
	res["hello"] = "world"
	ctx.JSON(res)
	return ctx.JSON(res)
}
