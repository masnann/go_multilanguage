package producthandler

import (
	"log"
	"multilanguage/constants"
	"multilanguage/handler"
	"multilanguage/helpers"
	"multilanguage/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	handler handler.Handler
}

func NewProductHandler(handler handler.Handler) ProductHandler {
	return ProductHandler{
		handler: handler,
	}
}

func (h ProductHandler) FindListProduct(ctx echo.Context) error {
	var result models.Response
	req := new(models.RequestListProduct)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	productList, err := h.handler.ProductService.FindListProduct(*req)
	if err != nil {
		log.Printf("Error FindListProduct: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_STRING, productList)
	return ctx.JSON(http.StatusOK, result)
}

func (h ProductHandler) CreateProduct(ctx echo.Context) error {
	var result models.Response
	req := new(models.ProductCreateRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	productID, err := h.handler.ProductService.CreateProduct(*req)
	if err != nil {
		log.Printf("Error CreateProduct: %v", err)
		result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_STRING, productID)
	return ctx.JSON(http.StatusCreated, result)
}
