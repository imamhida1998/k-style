package controller

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type transaksiController struct {
	transaksi usecase.Transaksi
}

func NewHandlerTransaksi(transaksi usecase.Transaksi) *transaksiController {
	return &transaksiController{transaksi}
}

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
