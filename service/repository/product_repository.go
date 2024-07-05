package repository

import (
	"fmt"
	"k-style/db"
	"k-style/service/model"
	"k-style/service/model/request"
)

type Product interface {
	CreateProduct(req *model.Product) error
	GetProductById(id string) (res model.Product, err error)
	UpdateProduct(userId, req *request.UpdateProduct) error
}

type productRepo struct {
}

func NewProductRepo() *productRepo {
	return &productRepo{}
}

func (p *productRepo) CreateProduct(req *model.Product) error {
	query := `
	insert
		product
	into
		(
		id,
		nama,
		kategori,
		harga,
		created_at,
		created_by)
	values
		(
		?,
		?,
		?,
		?,
		NOW(),
		?)`

	_, err := db.MySQL.Exec(query, req.Id, req.Nama, req.Kategori, req.Harga, req.CreatedBy)
	if err != nil {
		fmt.Print("err: ", err)
		return err
	}
	return nil
}

func (p *productRepo) GetProductById(id string) (res model.Product, err error) {
	query := `select id,nama,kategori,harga,created_at,created_by,updated_at,updated_by from product where id = ?`

	result, err := db.MySQL.Query(query, id)
	if err != nil {
		return res, err
	}

	for result.Next() {
		errx := result.Scan(
			&res.Id,
			&res.Nama,
			&res.Kategori,
			&res.Harga,
			&res.CreatedAt,
			&res.CreatedBy,
			&res.UpdatedAt,
			&res.UpdatedBy,
		)
		if err != nil {
			return res, errx
		}
	}
	return res, nil
}

func (p *productRepo) GetListProduct() (res model.Product, err error) {
	query := `select * from product`

	result, err := db.MySQL.Query(query)
	if err != nil {
		return res, err
	}

	for result.Next() {
		errx := result.Scan(
			&res.Id,
			&res.Nama,
			&res.Kategori,
			&res.Harga,
			&res.CreatedAt,
			&res.CreatedBy,
			&res.UpdatedAt,
			&res.UpdatedBy,
		)
		if err != nil {
			return res, errx
		}
	}
	return res, nil
}

func (p *productRepo) UpdateProduct(userId, req *request.UpdateProduct) error {
	query := `
		update
			product
		set
			nama = ?,
			kategori = ?,
			harga = ?,
			updated_at = now(),
			updated_by = ?
		where
			id = ?
			`

	if _, err := db.MySQL.Exec(query, req.Nama, req.Kategori, req.Harga, userId, req.ProductId); err != nil {
		return err
	}

	return nil
}
