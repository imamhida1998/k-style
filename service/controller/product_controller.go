package controller

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productController struct {
	productUsecase usecase.Product
}

func NewHandlerProduct(product usecase.Product) *productController {
	return &productController{product}
}

func (t *productController) CreateProduct(c echo.Context) error {
	var input request.CreateProduct

	currentUser := c.Get("CurrentUser").(model.User)

	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	err = t.productUsecase.CreateProduct(currentUser.Id, &input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	resp := echo.Map{
		"status":  201,
		"message": "success"}
	return c.JSON(http.StatusCreated, resp)

}
