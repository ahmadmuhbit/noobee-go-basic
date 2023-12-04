package auth

import "time"

type Auth struct {
	Id        int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(email, password string) Auth {
	return Auth{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a Auth) WithId(id int) Auth {
	a.Id = id
	return a
}
