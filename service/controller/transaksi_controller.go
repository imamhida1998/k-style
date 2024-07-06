package controller

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type transaksiController struct {
	transaksi usecase.Transaksi
}

func NewHandlerTransaksi(transaksi usecase.Transaksi) *transaksiController {
	return &transaksiController{transaksi}
}

// PaymentTransaksi handles example endpoint
// @Summary Get example
// @Description Get example
// @ID get-example
// @Produce json
// @Success 200 {string} string "ok"
// @Router /example [get]
func (t *transaksiController) PaymentTransaksi(c echo.Context) error {
	var input request.PaymentTransaksi

	currentUser := c.Get("CurrentUser").(model.User)
	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	err = t.transaksi.PaymentTransaksi(currentUser.Id, &input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}
	resp := echo.Map{
		"message": "success",
		"status":  200,
		"data": echo.Map{
			"statusPaid": "pending",
		},
	}

	return c.JSON(http.StatusBadRequest, resp)
}

// CancelTransaksi godoc
// @Summary Cancal Transaksi
// @Description Cancal Order pada product
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  data body request.PaymentTransaksi true "insert data"
// @Success 200 {object} interface{}
// @Router /api/users/payment/create [POST]
// @Tags Order Management
func (t *transaksiController) CancelTransaksi(c echo.Context) error {
	var input request.PaymentTransaksi

	currentUser := c.Get("CurrentUser").(model.User)
	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	err = t.transaksi.CancelTransaksi(input.TransaksiId, currentUser)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}
	resp := echo.Map{
		"message": "success",
		"status":  200,
		"data": echo.Map{
			"statusPaid": "pending",
		},
	}

	return c.JSON(http.StatusBadRequest, resp)
}
