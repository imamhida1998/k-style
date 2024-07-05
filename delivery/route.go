package delivery

import (
	"k-style/service/controller"
	"k-style/service/usecase"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

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

	api := e.Group("/api")

	api.POST("/register", handler.RegistrationDataUser)
	api.POST("/login", handler.Login)
	// api.GET("/product/:id", handler.Login)

	admin := api.Group("/admin")
	admin.Use(authMiddleware(auth, userUsecase))
	admin.GET("/", handler.DetailUser)
	// admin.DELETE("/delete", handler.Login)
	admin.PUT("/update", handler.UpdateUser)
	admin.POST("/delete-user", handler.DetailUser)

	productAdmin := admin.Group("/product")
	productAdmin.PUT("/update", handler.UpdateUser)
	productAdmin.POST("/create", handlerProduct.CreateProduct)
	// admin.GET("/list-payment", handler.DetailUser)

	user := api.Group("/users")
	user.Use(authMiddleware(auth, userUsecase))
	user.GET("/detail", handler.DetailUser)
	user.PUT("/update", handler.UpdateUser)
	// user.DELETE("/:id", handler.Login)
	// user.POST("/order", handler.Login)

	// product := user.Group("/product")
	// user.GET("/list", handler.DetailUser)

	payment := user.Group("/payment")
	payment.POST("/list", handlerTransaksi.PaymentTransaksi)
	payment.POST("/create", handlerTransaksi.PaymentTransaksi)
	payment.DELETE("/delete", handlerTransaksi.CancelTransaksi)

	tx := user.Group("/transaksi")
	tx.GET("/list", handlerProduct.CreateProduct)

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
