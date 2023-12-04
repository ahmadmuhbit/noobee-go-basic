package menu

import "time"

type Menu struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Desc      string    `json:"desc"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name, category, desc string, price int) Menu {
	return Menu{
		Name:      name,
		Category:  category,
		Desc:      desc,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m Menu) WithId(id int) Menu {
	m.Id = id
	return m
}
