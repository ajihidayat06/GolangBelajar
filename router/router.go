package router

import (
	"D/LatihanAji/GolangBelajar/model"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func ReadRequest(req interface{}, ctx *fiber.Ctx) (id int64, err error) {

	if ctx.Method() != "GET" {
		err = ctx.BodyParser(req)
		if err != nil {
			return
		}
	}
	fmt.Println(ctx.Params("id"))

	if ctx.Params("id") != "" {
		idInt, errs := ReadRequestParams(ctx)
		if errs != nil {
			err = errs
			return
		}
		id = int64(idInt)
	}

	return
}

func SetResponseJson(ctx *fiber.Ctx, resp model.Response) error {
	response := model.Response{
		StatusCode: resp.StatusCode,
		Message:    resp.Message,
		Data:       resp.Data,
	}
	return ctx.Status(resp.StatusCode).JSON(response)
}

func ReadRequestParams(ctx *fiber.Ctx) (id int64, err error) {
	idInt, errs := strconv.Atoi(ctx.Params("id"))
	if errs != nil {
		err = errs
		return
	}
	fmt.Println(id)
	id = int64(idInt)

	return
}
