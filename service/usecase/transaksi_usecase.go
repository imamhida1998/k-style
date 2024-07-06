package usecase

import (
	"errors"
	"fmt"
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/model/response"
	"k-style/service/repository"

	"github.com/google/uuid"
)

type Transaksi interface {
	CreateTransaksi(idUser string, req *request.CreateTransaksi) (res *response.CreateTransaksi, err error)
	PaymentTransaksi(userId string, req *request.PaymentTransaksi) error
	CancelTransaksi(transaksiId string, user model.User) error
	GetListPayment(UserId, Status string, page int, size int) (res model.Transaksi, err error)
}

type transaksiUsercase struct {
	tx      repository.TransaksiRepo
	product repository.Product
	user    repository.UserRepo
	auth    *JWTService
}

func NewTransaksiUsecase(tx repository.TransaksiRepo, product repository.Product, users repository.UserRepo, auth *JWTService) Transaksi {
	return &transaksiUsercase{tx, product, users, auth}
}

func (t *transaksiUsercase) CreateTransaksi(idUser string, req *request.CreateTransaksi) (res *response.CreateTransaksi, err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	product, err := t.product.GetProductById(req.ProductId)
	if err != nil {
		return nil, err
	}

	total := req.Quantity + product.Harga

	err = t.tx.CreateTransaksi(&model.Transaksi{
		Id:          id.String(),
		UserId:      idUser,
		NamaProduct: product.Nama,
		Status:      "unpaid",
		Total:       total,
	})
	if err != nil {
		return nil, err
	}

	resp := &response.CreateTransaksi{
		TransaksiId: id.String(),
		Total:       total,
	}
	return resp, nil
}

func (t *transaksiUsercase) PaymentTransaksi(userId string, req *request.PaymentTransaksi) error {

	transaki, err := t.tx.GetTransaksiById(req.TransaksiId)
	if err != nil {
		return err
	}

	if req.Amount < transaki.Total {
		count := transaki.Total - req.Amount
		message := fmt.Sprintf("transaksi anda kurang %d", count)
		return errors.New(message)
	}

	err = t.tx.UpdateStatusTransaksiById(req.TransaksiId, "", "pending")
	if err != nil {
		return err
	}
	return nil

}

func (t *transaksiUsercase) CancelTransaksi(transaksiId string, user model.User) error {
	transaki, err := t.tx.GetTransaksiById(transaksiId)
	if err != nil {
		return err
	}

	if transaki.UserId != user.Id {
		if user.Role == "Admin" {
			err = t.tx.UpdateStatusTransaksiById(transaksiId, user.Id, "cancelled")
			if err != nil {
				return err
			}
		} else {
			return errors.New("tidak dapat cancel transaksi orang lain !")
		}
	}
	err = t.tx.UpdateStatusTransaksiById(transaksiId, user.Id, "cancelled")
	if err != nil {
		return err
	}

	return nil
}

func (t *transaksiUsercase) GetListPayment(UserId, Status string, page int, size int) (res model.Transaksi, err error) {

	res, err = t.tx.GetTransaksiByStatusUserId(page, size, Status, UserId)
	if err != nil {
		return res, err
	}

	return res, nil

}
