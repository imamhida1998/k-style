package repository

import (
	"database/sql"
	"errors"
	"k-style/db"
	"k-style/service/model"
	"k-style/service/model/request"
	"log"
)

type UserRepo interface {
	Registrasi(req *model.User) error
	Login(input *request.Login) (res model.User, err error)
	UpdateDataUsers(id string, data *model.User) error
	GetUsersByEmail(email string) (res model.User, err error)
	GetUsersById(id string) (res model.User, err error)
}

type repoUser struct {
}

func NewRepoUser() *repoUser {
	return &repoUser{}
}

func (r *repoUser) Registrasi(req *model.User) error {
	query := `
insert 
			into 
		users 
			(
			id,
			fullname,
			username,
			role,
			password,
			email,
			created_at) 
		values 
			(?,
			 ?,
			 ?,
			 ?,
			 ?,
			 ?,
			NOW())`
	_, err := db.MySQL.Exec(query, req.Id, req.Fullname, req.Username, req.Role, req.Password, req.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repoUser) Login(input *request.Login) (res model.User, err error) {
	query := `select id,fullname,username,role,password,email,created_at from users where username = ? and password = ?`
	rows, err := db.MySQL.Query(query, input.Username, input.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		errx := rows.Scan(
			&res.Id,
			&res.Fullname,
			&res.Username,
			&res.Role,
			&res.Password,
			&res.Email,
			&res.CreatedAt)
		if errx != nil {
			return res, errx
		}

	}
	return res, nil

}

func (r *repoUser) GetUsersByEmail(email string) (res model.User, err error) {
	query := `select id,fullname,username,password,role,email,created_at  from users where email = ?`
	rows, err := db.MySQL.Query(query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		errx := rows.Scan(
			&res.Id,
			&res.Fullname,
			&res.Username,
			&res.Password,
			&res.Role,
			&res.Email,
			&res.CreatedAt)
		if errx != nil {
			return res, errx
		}

	}
	return res, nil

}

func (r *repoUser) UpdateDataUsers(id string, data *model.User) error {
	query := `
		update
			users
		set
		fullname = ?,
		username = ?,
		password = ?,
		email = ?,
		updated_at = NOW()
	where
		id = ?
		`

	_, err := db.MySQL.Exec(query, data.Fullname, data.Username, data.Password, data.Email, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repoUser) GetUsersById(id string) (res model.User, err error) {
	query := `select id,fullname,username,password,role,email,created_at  from users where id = ?`
	rows, err := db.MySQL.Query(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		errx := rows.Scan(
			&res.Id,
			&res.Fullname,
			&res.Username,
			&res.Password,
			&res.Role,
			&res.Email,
			&res.CreatedAt)
		if errx != nil {
			return res, errx
		}

	}
	return res, nil

}
