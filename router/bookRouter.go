package router

import (
	"D/LatihanAji/GolangBelajar/model"
	"D/LatihanAji/GolangBelajar/usecase"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BookRouter struct {
	BookUseCase usecase.BookUseCase
}

func NewBookRouter(bookUseCase *usecase.BookUseCase) BookRouter {
	return BookRouter{BookUseCase: *bookUseCase}
}

func (r *BookRouter) Route(app *fiber.App) {
	app.Post("/api/book", r.InsertBook)
	app.Get("/api/book", r.GetAll)
	app.Get("/api/book/:id", r.GetDetail)
	app.Put("/api/book/:id", r.Update)
	app.Delete("/api/book/:id", r.Delete)
}

func (r *BookRouter) InsertBook(ctx *fiber.Ctx) error {
	var book model.Book
	_, err := ReadRequest(&book, ctx)
	fmt.Println(err)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// validation
	err = Validator.Struct(book)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// save data
	response, err := r.BookUseCase.Create(book)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success insert book",
		Data:       response,
	})
}

func (r *BookRouter) GetAll(ctx *fiber.Ctx) error {
	responses := r.BookUseCase.List()

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success get list books",
		Data:       responses,
	})
}

func (r *BookRouter) GetDetail(ctx *fiber.Ctx) error {
	id, err := ReadRequestParams(ctx)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	response, err := r.BookUseCase.GetDetail(id)
	if err != nil {
		var statusCode int
		if err.Error() != "book with this id is empty" {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusBadRequest
		}

		return SetResponseJson(ctx, model.Response{
			StatusCode: statusCode,
			Message:    err.Error(),
		})
	}

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success get setail book",
		Data:       response,
	})
}

func (r *BookRouter) Update(ctx *fiber.Ctx) error {
	var book model.Book
	id, err := ReadRequest(&book, ctx)
	fmt.Println(err)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// validation
	err = Validator.Struct(book)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// save data
	_, err = r.BookUseCase.Update(&book, id)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success update book",
		Data:       book,
	})
}

func (r *BookRouter) Delete(ctx *fiber.Ctx) error {
	id, err := ReadRequestParams(ctx)
	if err != nil {
		return SetResponseJson(ctx, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	err = r.BookUseCase.Delete(id)
	if err != nil {
		var statusCode int
		if err.Error() != "book with this id is empty" {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusBadRequest
		}

		return SetResponseJson(ctx, model.Response{
			StatusCode: statusCode,
			Message:    err.Error(),
		})
	}

	return SetResponseJson(ctx, model.Response{
		StatusCode: http.StatusOK,
		Message:    "success get delete book",
		Data:       "",
	})
}
