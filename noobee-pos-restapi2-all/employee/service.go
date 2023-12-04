package employee

import (
	"log"
)

type RepositoryInterface interface {
	GetAll() (employees []Employee, err error)
	Create(emp Employee) (err error)
	UpdateById(emp Employee, id int) (err error)
	DeleteById(id int) (err error)
	GetById(id int) (emp Employee, err error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) FindAll() (employees []Employee, err error) {
	employees, err = s.repo.GetAll()
	if err != nil {
		log.Println("error when try to GetAll employee with error", err.Error())
		return
	}

	return
}

func (s Service) Create(req Employee) (err error) {
	err = s.repo.Create(req)
	if err != nil {
		log.Println("error when try to Create new employee with error", err.Error())
		return
	}

	return
}

func (s Service) UpdateById(req Employee, id int) (err error) {
	err = s.repo.UpdateById(req, id)
	if err != nil {
		log.Println("error when try to UpdateById employee with error", err.Error())
		return
	}

	return
}

func (s Service) DeleteById(id int) (err error) {
	// err = s.repo.DeleteById(id)
	// if err != nil {
	// 	log.Println("error when try to Delete employee with error", err.Error())
	// 	return
	// }

	// return

	// Validate is exists
	_, err = s.repo.GetById(id)
	if err != nil {
		return
	}

	return s.repo.DeleteById(id)
}

func (s Service) GetById(id int) (emp Employee, err error) {
	emp, err = s.repo.GetById(id)
	if err != nil {
		log.Println("error when try to get employee with error", err.Error())
		return
	}

	return
}
