package controller

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transaksiController struct {
	transaksi usecase.Transaksi
}

func NewHandlerTransaksi(transaksi usecase.Transaksi) *transaksiController {
	return &transaksiController{transaksi}
}

// PaymentTransaksi godoc
// @Summary Payment Order
// @Description Payment Order pada product
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  data body request.PaymentTransaksi true "insert data"
// @Success 200 {object} interface{}
// @Router /api/users/payment/update [PUT]
// @Tags Customer Management
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

	return c.JSON(http.StatusOK, resp)
}

// CancelTransaksi godoc
// @Summary Cancal Transaksi
// @Description Cancal Order pada product
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  data body request.CancelTransaksi true "insert data"
// @Success 200 {object} interface{}
// @Router /api/users/payment/delete [DELETE]
// @Tags Customer Management
func (t *transaksiController) CancelTransaksi(c echo.Context) error {
	var input request.CancelTransaksi

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
			"statusPaid": "cancelled",
		},
	}

	return c.JSON(http.StatusOK, resp)
}

// CancelTransaksi godoc
// @Summary Get List Transaksi By Status
// @Description Cancal Order pada product
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  page query int true "Page number" default(0)
// @Param  size query int true "Items per page" default(0)
// @Param  status query string true "Filter status"
// @Param  user_id query string false "Filter userId"
// @Success 200 {object} interface{}
// @Router /api/users/transaksi/list [GET]
// @Tags Customer Management
func (t *transaksiController) GetListPayment(c echo.Context) error {
	// var input request.PaymentTransaksi

	page := c.QueryParam("page")
	size := c.QueryParam("size")
	userId := c.QueryParam("user_id")
	status := c.QueryParam("status")
	sizePage, _ := strconv.Atoi(size)
	noPage, _ := strconv.Atoi(page)
	currentUser := c.Get("CurrentUser").(model.User)

	if userId == "" {
		userId = currentUser.Id
	}

	if currentUser.Role != "Admin" {
		res, err := t.transaksi.GetListPayment(userId, status, noPage, sizePage)
		if err != nil {
			MessageError := echo.Map{"errors": err.Error()}
			return c.JSON(http.StatusBadRequest, MessageError)
		}
		resp := echo.Map{
			"message": "success",
			"status":  200,
			"data":    res,
		}

		return c.JSON(http.StatusOK, resp)
	} else {
		res, err := t.transaksi.GetListPayment(currentUser.Id, status, noPage, sizePage)
		if err != nil {
			MessageError := echo.Map{"errors": err.Error()}
			return c.JSON(http.StatusBadRequest, MessageError)
		}
		resp := echo.Map{
			"message": "success",
			"status":  200,
			"data":    res,
		}

		return c.JSON(http.StatusOK, resp)
	}

}

// PaymentTransaksi godoc
// @Summary Order Product
// @Description Create Transaksi
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  data body request.CreateTransaksi true "insert data"
// @Success 200 {object} interface{}
// @Router /api/users/payment/create [POST]
// @Tags Customer Management
func (t *transaksiController) CreateTransaksi(c echo.Context) error {
	var input request.CreateTransaksi

	currentUser := c.Get("CurrentUser").(model.User)
	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	res, err := t.transaksi.CreateTransaksi(currentUser.Id, &input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}
	resp := echo.Map{
		"message": "success",
		"status":  200,
		"data":    res,
	}

	return c.JSON(http.StatusOK, resp)
}

// AcceptTransaksi godoc
// @Summary Accept Payment Transaction
// @Description Hanya Role Admin yang dapat menggunakan akses ini
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  data body request.AcceptTransaksi true "insert data"
// @Success 200 {object} interface{}
// @Router /api/users/payment/accept-payment [PUT]
// @Tags Customer Management
func (t *transaksiController) AcceptTransaksi(c echo.Context) error {
	var input request.AcceptTransaksi

	currentUser := c.Get("CurrentUser").(model.User)
	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	res, err := t.transaksi.AcceptPayment(input.TransaksiId, currentUser.Id)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}
	resp := echo.Map{
		"message": "success",
		"status":  200,
		"data":    res,
	}

	return c.JSON(http.StatusOK, resp)
}
