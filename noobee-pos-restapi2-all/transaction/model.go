package transaction

import "time"

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Menu struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Transaction struct {
	Id         int       `json:"id"`
	MenuId     int       `json:"menu_id"`
	EmployeeId int       `json:"employee_id"`
	Quantity   int       `json:"quantity"`
	Total      int       `json:"total"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// For join
	Employee Employee `json:"employee"`
	Menu     Menu     `json:"menu"`
}

func NewFromPayRequest(req PayRequest) Transaction {
	return Transaction{
		MenuId:     req.MenuId,
		EmployeeId: req.EmployeeId,
		Quantity:   req.Quantity,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

}

func (t *Transaction) SetTotal(price int) {
	t.Total = t.Quantity * price
}
