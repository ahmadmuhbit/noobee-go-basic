package menu

type CreateMenuRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	Price    int    `json:"price"`
}
