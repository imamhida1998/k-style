package delivery

import (
	"k-style/service/controller"
	"k-style/service/usecase"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	// _ "github.com/swaggo/echo-swagger/example/docs"
	_ "k-style/docs"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func Route(
	e *echo.Echo,
	transaksi usecase.Transaksi,
	productUsecase usecase.Product,
	userUsecase usecase.Users,
	auth usecase.Auth,

) {
	handler := controller.NewHandlerUser(userUsecase)
	handlerProduct := controller.NewHandlerProduct(productUsecase)
	handlerTransaksi := controller.NewHandlerTransaksi(transaksi)
	e.Use(middleware.Recover())
	api := e.Group("/api")

	api.POST("/register", handler.RegistrationDataUser)
	api.POST("/login", handler.Login)
	api.GET("/swagger/*any", echoSwagger.WrapHandler)
	// api.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		if strings.Contains(c.Request().URL.Path, "swagger") {
	// 			return true
	// 		}
	// 		return false
	// 	},
	// }))
	// api.GET("/product/:id", handler.Login)

	admin := api.Group("/admin")
	admin.Use(authMiddleware(auth, userUsecase))
	admin.GET("/", handler.DetailUser)
	// admin.DELETE("/delete", handler.Login)
	admin.PUT("/update", handler.UpdateUser)
	admin.DELETE("/delete-user", handler.DetailUser)
	admin.GET("/list/payment", handlerTransaksi.GetListPayment)

	productAdmin := admin.Group("/product")
	productAdmin.PUT("/update", handlerProduct.UpdateProduct)
	productAdmin.POST("/create", handlerProduct.CreateProduct)
	productAdmin.GET("/list", handlerProduct.GetListProduct)
	productAdmin.DELETE("/delete", handlerProduct.DeleteProduct)
	// admin.GET("/list-payment", handler.DetailUser)

	user := api.Group("/users")
	user.Use(authMiddleware(auth, userUsecase))
	user.GET("/detail", handler.DetailUser)
	user.PUT("/update", handler.UpdateUser)
	admin.GET("/detail-payment", handlerTransaksi.GetListPayment)
	// user.DELETE("/:id", handler.Login)
	// user.POST("/order", handler.Login)

	// product := user.Group("/product")
	// user.GET("/list", handler.DetailUser)

	payment := user.Group("/payment")
	payment.POST("/create", handlerTransaksi.CreateTransaksi)
	payment.PUT("/update", handlerTransaksi.PaymentTransaksi)
	payment.DELETE("/delete", handlerTransaksi.CancelTransaksi)
	payment.PUT("/accept-payment", handlerTransaksi.AcceptTransaksi)
	payment.GET("/transaksi-detail", handlerTransaksi.GetTransaksiId)

	tx := user.Group("/transaksi")
	tx.GET("/list", handlerTransaksi.GetListPayment)

}

func authMiddleware(authService usecase.Auth, userService usecase.Users) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.Contains(authHeader, "Bearer") {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			var tokenString string
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}
			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			Email := (claim["email"].(string))

			user, err := userService.GetDetailUserByEmail(Email)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			c.Set("CurrentUser", user)

			return next(c)
		}
	}
}
