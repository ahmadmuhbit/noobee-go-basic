package employee

import "time"

type Employee struct {
	Id        int       `json:"id"`
	NIP       string    `json:"nip"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(nip, name, address string) Employee {
	return Employee{
		NIP:       nip,
		Name:      name,
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
