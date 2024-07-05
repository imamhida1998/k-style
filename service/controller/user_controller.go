package controller

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type userController struct {
	userUsecase usecase.Users
}

func NewHandlerUser(user usecase.Users) *userController {
	return &userController{user}
}

func (h *userController) RegistrationDataUser(c echo.Context) error {
	var input request.Register

	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}

		return c.JSON(http.StatusBadRequest, MessageError)
	}

	err = h.userUsecase.Register(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err}
		c.JSON(http.StatusInternalServerError, MessageError)
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	return c.JSON(http.StatusOK, "succes")

}

func (h *userController) Login(c echo.Context) error {
	var input request.Login

	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}

		return c.JSON(http.StatusBadRequest, MessageError)
	}

	res, err := h.userUsecase.Login(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, MessageError)
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	resp := echo.Map{
		"message": "success",
		"status":  200,
		"data":    res}

	return c.JSON(http.StatusOK, resp)

}

func (h *userController) DetailUser(c echo.Context) error {

	currentUser := c.Get("CurrentUser").(model.User)

	resp := echo.Map{
		"message": "success",
		"status":  200,
		"data":    currentUser}

	return c.JSON(http.StatusOK, resp)

}

func (h *userController) UpdateUser(c echo.Context) error {
	var input request.UpdateUser
	currentUser := c.Get("CurrentUser").(model.User)

	err := c.Bind(&input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}

		return c.JSON(http.StatusBadRequest, MessageError)
	}

	err = h.userUsecase.UpdateUser(&currentUser, &input)
	if err != nil {
		MessageError := echo.Map{"errors": err.Error()}
		return c.JSON(http.StatusBadRequest, MessageError)
	}

	resp := echo.Map{
		"message": "success",
		"status":  200,
	}

	return c.JSON(http.StatusOK, resp)

}
