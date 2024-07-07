package controller

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productController struct {
	productUsecase usecase.Product
}

func NewHandlerProduct(product usecase.Product) *productController {
	return &productController{product}
}

// CreateProduct godoc
// @Summary Create Product
// @Description Membuat Product yang akan dijual
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  data body request.CreateProduct true "insert data"
// @Success 200 {object} interface{}
// @Router /api/admin/product/create [POST]
// @Tags Order Management
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

// UpdateProduct godoc
// @Summary Update Product
// @Description Merubah Product yang akan dijual
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  data body request.CreateProduct true "insert data"
// @Success 200 {object} interface{}
// @Router /api/admin/product/update [PUT]
// @Tags Order Management
func (t *productController) UpdateProduct(c echo.Context) error {
	var input request.UpdateProduct

	currentUser := c.Get("CurrentUser").(model.User)
	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	res, err := t.productUsecase.GetListProducById(input.ProductId)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	if res.CreatedBy != currentUser.Id {
		if currentUser.Role == "Admin" {
			err = t.productUsecase.UpdateProduct(currentUser.Id, &input)
			if err != nil {
				MessageError := echo.Map{"errors": err.Error()}
				return c.JSON(http.StatusBadRequest, MessageError)
			}

			resp := echo.Map{
				"status":  201,
				"message": "success"}
			return c.JSON(http.StatusCreated, resp)
		}
		MessageError := echo.Map{"errors": "tidak dapat mendapatkan akses merubah product user lain"}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	err = t.productUsecase.UpdateProduct(currentUser.Id, &input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	resp := echo.Map{
		"status":  201,
		"message": "success"}
	return c.JSON(http.StatusCreated, resp)

}

// CancelTransaksi godoc
// @Summary List Product
// @Description List Data Product
// @Accept  application/json
// @Security BearerAuth
// @Produce  json
// @Param  page query int true "Page number" default(0)
// @Param  size query int true "Items per page" default(0)
// @Param  kategori query string false "Filter kategori"
// @Success 200 {object} interface{}
// @Router /api/admin/product/list [GET]
// @Tags Order Management
func (t *productController) GetListProduct(c echo.Context) error {

	page := c.QueryParam("page")
	size := c.QueryParam("size")
	kategori := c.QueryParam("kategori")
	sizePage, _ := strconv.Atoi(size)
	noPage, _ := strconv.Atoi(page)
	res, err := t.productUsecase.GetListProductWithPageSize(kategori, noPage, sizePage)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	resp := echo.Map{
		"status":  201,
		"message": "success",
		"data":    res}
	return c.JSON(http.StatusCreated, resp)

}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Menghapus Product
// @Accept  application/json
// @Produce  json
// @Security BearerAuth
// @Param  data body request.UpdateProduct true "insert data"
// @Success 200 {object} interface{}
// @Router /api/admin/product/delete [DELETE]
// @Tags Order Management
func (t *productController) DeleteProduct(c echo.Context) error {
	var input request.DeleteProduct

	currentUser := c.Get("CurrentUser").(model.User)
	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	if currentUser.Role != "Admin" {
		MessageError := echo.Map{"errors": "tidak dapat mendapatkan akses delete product"}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	err = t.productUsecase.DeleteProduct(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	resp := echo.Map{
		"status":  201,
		"message": "success"}
	return c.JSON(http.StatusCreated, resp)

}
