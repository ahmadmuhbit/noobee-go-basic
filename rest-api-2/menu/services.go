package menu

import "log"

type RepositoryInterface interface {
	Save(menu Menu) (err error)
	GetAll() (menus []Menu, err error)
	GetById(id int) (menu Menu, err error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Save(menu Menu) (err error) {
	err = s.repo.Save(menu)
	if err != nil {
		log.Println("error when try to save menu with error", err.Error())
		return
	}

	return
}

func (s Service) GetAll() (menus []Menu, err error) {
	menus, err = s.repo.GetAll()
	if err != nil {
		log.Println("error when try to get all menu with error", err.Error())
		return
	}

	return
}

func (s Service) GetById(id int) (menu Menu, err error) {
	menu, err = s.repo.GetById(id)
	if err != nil {
		log.Println("error when try to get menu with error", err.Error())
		return
	}

	return
}
