package auth

import (
	"database/sql"
	"fmt"
	"log"
	"nobee-pos/utility"
)

type RepositoryInterface interface {
	Create(auth Auth) (err error)
	GetByEmail(email string) (auth Auth, err error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Create(auth Auth) (err error) {
	auth.Password, err = utility.Hash(auth.Password)
	if err != nil {
		log.Println("error when try to hash password with error", err.Error())
		return
	}

	err = s.repo.Create(auth)
	if err != nil {
		log.Println("error when try to create auth with error", err.Error())
		return
	}
	return
}

func (s Service) Login(req Auth) (auth Auth, err error) {
	auth, err = s.repo.GetByEmail(req.Email)
	if err != nil {
		log.Println("error when try to create auth with error", err.Error())
		if err == sql.ErrNoRows {
			err = fmt.Errorf("username or password not found")
			return
		}
		return
	}

	err = utility.Verify(auth.Password, req.Password)
	if err != nil {
		log.Println("error when try to create auth with error", err.Error())
		err = fmt.Errorf("Username or password not found")
		return
	}
	return
}
