package menu

import "database/sql"

type MemoryRepo struct{}

var data = []Menu{}

func NewMemoryRepo() MemoryRepo {
	return MemoryRepo{}
}

func (m MemoryRepo) GetAll() (menus []Menu, err error) {
	menus = data
	return
}

func (m MemoryRepo) GetById(id int) (menu Menu, err error) {
	for _, d := range data {
		if d.Id == id {
			return d, nil
		}
	}

	err = sql.ErrNoRows
	return
}

func (m MemoryRepo) Save(menu Menu) (err error) {
	menu.Id = len(data) + 1
	data = append(data, menu)
	return
}
