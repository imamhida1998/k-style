package main

import (
	"k-style/db"
	"k-style/delivery"
	"k-style/service/repository"
	"k-style/service/usecase"
	"k-style/util"

	"github.com/labstack/echo"
)

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

	e.Start(":8000")

}
