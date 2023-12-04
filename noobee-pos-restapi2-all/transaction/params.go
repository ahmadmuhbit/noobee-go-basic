package transaction

type PayRequest struct {
	MenuId     int `json:"menu_id"`
	EmployeeId int `json:"employee_id"`
	Quantity   int `json:"quantity"`
}
