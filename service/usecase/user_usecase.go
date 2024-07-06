package usecase

import (
	"errors"
	"k-style/service/model"
	"k-style/service/model/request"
	"k-style/service/model/response"
	"k-style/service/repository"
	"k-style/util"

	"github.com/google/uuid"
)

type Users interface {
	Register(params *request.Register) error
	Login(params *request.Login) (res *response.Login, err error)
	GetDetailUserByEmail(email string) (res model.User, err error)
	UpdateUser(user *model.User, reqUpdate *request.UpdateUser) error
}

type userUsercase struct {
	user repository.UserRepo
	auth *JWTService
}

func NewUserUsecase(users repository.UserRepo, auth *JWTService) Users {
	return &userUsercase{users, auth}
}

func (u *userUsercase) Register(params *request.Register) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	role, err := util.GetRoleByTypeId(params.RoleId)
	if err != nil {
		return err
	}

	return u.user.Registrasi(&model.User{
		Id:       id.String(),
		Username: params.Username,
		Fullname: params.Fullname,
		Role:     role,
		Email:    params.Email,
		Password: params.Password,
	})
}

func (u *userUsercase) Login(params *request.Login) (res *response.Login, err error) {

	user, err := u.user.Login(params)
	if err != nil {
		return
	}
	if user.Email == "" {
		return nil, errors.New("Login gagal")
	}
	token, err := u.auth.GenerateTokenJWT(user.Email)
	if err != nil {
		return
	}

	resp := &response.Login{
		AccessToken: "Bearer " + token,
	}

	return resp, nil
}

func (u *userUsercase) GetDetailUserByEmail(email string) (res model.User, err error) {

	res, err = u.user.GetUsersByEmail(email)
	if err != nil {
		return
	}

	return res, nil
}

func (u *userUsercase) UpdateUser(user *model.User, reqUpdate *request.UpdateUser) error {

	if user.Fullname != reqUpdate.Fullname {
		user.Fullname = reqUpdate.Fullname
	}
	if user.Email != reqUpdate.Email {
		user.Email = reqUpdate.Email
	}

	if user.Role != reqUpdate.Role {
		if user.Role == "Admin" {
			user.Role = reqUpdate.Role
		} else {
			if user.Role == "customer" {
				if reqUpdate.Role == "Admin" {
					return errors.New("Anda tidak dapat mengganti role Admin")
				}
			}

		}

	}

	if user.Password != reqUpdate.Password {
		user.Password = reqUpdate.Password
	}

	err := u.user.UpdateDataUsers(user.Id, user)
	if err != nil {
		return err
	}

	return nil
}
