package employee

type CreateRequest struct {
	NIP     string `json:"nip"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UpdateRequest struct {
	NIP     string `json:"nip"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
