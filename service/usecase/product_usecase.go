package usecase

import (
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/repository"

	"github.com/google/uuid"
)

type Product interface {
	CreateProduct(idUser string, req *request.CreateProduct) error
	UpdateProduct(userId string, updateProduct *request.UpdateProduct) error
	GetListProducById(Id string) (res model.Product, err error)
	GetListProductWithPageSize(kategori string, page, size int) (res []model.Product, err error)
	DeleteProduct(req *request.DeleteProduct) error
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

func (p *productUsercase) UpdateProduct(userId string, updateProduct *request.UpdateProduct) error {

	product, err := p.product.GetProductById(updateProduct.ProductId)
	if err != nil {
		return err
	}

	if product.Harga != updateProduct.Harga {
		product.Harga = updateProduct.Harga
	}
	if product.Nama != updateProduct.Nama {
		product.Nama = updateProduct.Nama
	}

	if product.Kategori != updateProduct.Kategori {
		product.Kategori = updateProduct.Kategori
	}

	err = p.product.UpdateProduct(userId, &product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productUsercase) GetListProductWithPageSize(kategori string, page, size int) (res []model.Product, err error) {
	if kategori == "" {
		res, err = p.product.GetProductWithPageSize(page, size)
		if err != nil {
			return res, nil
		}
	} else {
		res, err = p.product.GetKategoriProductWithPageSize(kategori, page, size)
		if err != nil {
			return res, nil
		}
	}

	return res, nil
}

func (p *productUsercase) DeleteProduct(req *request.DeleteProduct) error {
	return p.product.DeleteProduct(req)
}

func (p *productUsercase) GetListProducById(Id string) (res model.Product, err error) {
	return p.product.GetProductById(Id)
}
