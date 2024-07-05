package usecase

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/repository"

	"github.com/google/uuid"
)

type Product interface {
	CreateProduct(idUser string, req *request.CreateProduct) error
}

type productUsercase struct {
	product repository.Product
	user    repository.UserRepo
	auth    *JWTService
}

func NewProductUsecase(product repository.Product, users repository.UserRepo, auth *JWTService) Product {
	return &productUsercase{product, users, auth}
}

func (p *productUsercase) CreateProduct(idUser string, req *request.CreateProduct) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	return p.product.CreateProduct(&model.Product{
		Id:        id.String(),
		Nama:      req.Nama,
		Harga:     req.Harga,
		Kategori:  req.Kategori,
		CreatedBy: idUser,
	})

}

// func (p *productUsercase) UpdateProduct(userId string, updateProduct *request.UpdateProduct) error {
// 	product, err := p.product.GetProductById(updateProduct.ProductId)
// 	if err != nil {
// 		return err
// 	}

// 	p.product.UpdateProduct()

// }
