package main

import (
	"k-style/db"
	"k-style/delivery"
	"k-style/service/repository"
	"k-style/service/usecase"
	"k-style/util"

	_ "k-style/docs"

	"github.com/labstack/echo/v4"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://example.com/contact/
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()
	db.InitDB()
	util.HandleCors(e)
	user := repository.NewRepoUser()
	product := repository.NewProductRepo()
	tx := repository.NewRepoTransaksi()

	jwtAuth := usecase.NewJWTService()
	userUsercase := usecase.NewUserUsecase(user, jwtAuth)
	productUsecase := usecase.NewProductUsecase(product, user, jwtAuth)
	transaksiUsecase := usecase.NewTransaksiUsecase(tx, product, user, jwtAuth)

	delivery.Route(e, transaksiUsecase, productUsecase, userUsercase, jwtAuth)

	err := e.Start("0.0.0.0:3000")
	if err != nil {
		return
	}

}
