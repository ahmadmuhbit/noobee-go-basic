package transaction

import (
	"database/sql"
	"errors"
	"log"
)

type RepositoryInterface interface {
	FindAll() (trxs []Transaction, err error)
	FindById(id int) (trx Transaction, err error)
	Create(trx Transaction) (err error)

	MenuRepositoryInterface
	EmployeeRepositoryInterface
}

type MenuRepositoryInterface interface {
	GetMenuById(id int) (menu Menu, err error)
}

type EmployeeRepositoryInterface interface {
	GetEmployeeById(id int) (emp Employee, err error)
}

type Service struct {
	repo RepositoryInterface
}

func NewServcie(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) ListAll() (trxs []Transaction, err error) {
	trxs, err = s.repo.FindAll()
	return
}

func (s Service) Detail(trxId int) (trx Transaction, err error) {
	trx, err = s.repo.FindById(trxId)
	return
}

func (s Service) Pay(req Transaction) (err error) {
	_, err = s.repo.GetEmployeeById(req.EmployeeId)
	if err != nil {
		log.Println("error when try to GetEmployeeById with emp_id:", req.Employee.Id, "and error:", err.Error())
		if err == sql.ErrNoRows {
			err = errors.New("menu not found")
		}
		return
	}
	menu, err := s.repo.GetMenuById(req.MenuId)
	if err != nil {
		log.Println("error when try to GetMenuById with menu_id:", req.MenuId, "and error", err.Error())
		if err == sql.ErrNoRows {
			err = errors.New("menu not found")
		}
		return
	}
	req.SetTotal(menu.Price)

	err = s.repo.Create(req)
	return
}
