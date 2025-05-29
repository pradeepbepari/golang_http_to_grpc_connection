package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	Contact   string `json:"contact"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
