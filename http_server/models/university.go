package models

type University struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Contact   string `json:"contact"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
