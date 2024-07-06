package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"k-style/db"
	"k-style/service/model"
)

type TransaksiRepo interface {
	CreateTransaksi(req *model.Transaksi) error
	GetTransaksiById(id string) (res model.Transaksi, err error)
	UpdateStatusTransaksiById(id, userId, status string) error
	GetTransaksiByStatus(status string) (res model.Transaksi, err error)
	GetAllTransaksi() (res model.Transaksi, err error)
	GetTransaksiByStatusUserId(page int, size int, status, userId string) (res []model.Transaksi, err error)
}
type repoTransaksi struct {
}

func NewRepoTransaksi() *repoTransaksi {
	return &repoTransaksi{}
}

func (t *repoTransaksi) CreateTransaksi(req *model.Transaksi) error {
	query := `
	insert
		into
			transaksi
		(
			id,
			user_id,
			product_id,
			nama_product,
			status,
			total,
			created_at)
		values
			(
			?,
			?,
			?,
			?,
			?,
			?,
			NOW())`

	_, err := db.MySQL.Exec(query, req.Id, req.UserId, req.ProductId, req.NamaProduct, req.Status, req.Total)
	if err != nil {
		fmt.Print("err: ", err.Error())
		return err
	}
	return nil

}

func (t *repoTransaksi) UpdateStatusTransaksiById(id, userId, status string) error {
	query := `
			update 
				transaksi
			set
				status = ?,
				updated_by = ?,
				updated_at = now()
			where
				id = ?`
	if _, err := db.MySQL.Exec(query, status, userId, id); err != nil {
		return err
	}
	return nil
}

func (t *repoTransaksi) GetTransaksiById(id string) (res model.Transaksi, err error) {
	query := `
	select 
			id,
			user_id,
			product_id,
			nama_product,
			status,
			total,
			created_at
	from
			transaksi
	where
			id = ?`

	result, err := db.MySQL.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("tidak dapat menemukan transaksi")
		}
		return res, err
	}

	for result.Next() {
		err = result.Scan(
			&res.Id,
			&res.UserId,
			&res.ProductId,
			&res.NamaProduct,
			&res.Status,
			&res.Total,
			&res.CreatedAt,
		)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func (t *repoTransaksi) GetTransaksiByStatus(status string) (res model.Transaksi, err error) {
	query := `
	select 
			id,
			user_id,
			nama_product,
			status,
			total,
			created_at
	from
			transaksi
	where
			status = ?
	order by
			created_at
	desc`

	result, err := db.MySQL.Query(query, status)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("tidak ada transaksi")
		}
		return res, err
	}

	for result.Next() {
		err = result.Scan(
			&res.Id,
			&res.UserId,
			&res.NamaProduct,
			&res.Status,
			&res.Total,
			&res.CreatedAt,
		)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func (t *repoTransaksi) GetAllTransaksi() (res model.Transaksi, err error) {
	query := `
	select 
			id,
			user_id,
			nama_product,
			status,
			total,
			created_at
	from
			transaksi
	order by
			created_at
	desc`

	result, err := db.MySQL.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("tidak ada transaksi")
		}
		return res, err
	}

	for result.Next() {
		err = result.Scan(
			&res.Id,
			&res.UserId,
			&res.NamaProduct,
			&res.Status,
			&res.Total,
			&res.CreatedAt,
		)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func (t *repoTransaksi) GetTransaksiByStatusUserId(page int, size int, status, userId string) (res []model.Transaksi, err error) {
	// Hitung offset berdasarkan halaman
	NoPage := (page - 1) * size

	// Query SQL
	query := `
		SELECT 
			id,
			user_id,
			nama_product,
			status,
			total,
			created_at
		FROM
			transaksi
		WHERE
			status = ?
			AND user_id = ?
		ORDER BY
			created_at DESC
		LIMIT ?
		OFFSET ?`

	// Eksekusi query
	result, err := db.MySQL.Query(query, status, userId, size, NoPage)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("tidak ada transaksi")
		}
		return res, err
	}
	defer result.Close()

	// Iterasi hasil dan scan ke dalam struct
	for result.Next() {
		var transaksi model.Transaksi
		err = result.Scan(
			&transaksi.Id,
			&transaksi.UserId,
			&transaksi.NamaProduct,
			&transaksi.Status,
			&transaksi.Total,
			&transaksi.CreatedAt,
		)
		if err != nil {
			return res, err
		}
		res = append(res, transaksi)
	}

	return res, nil
}
