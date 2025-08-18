package models

type Role string

const (
	AdminRole      Role = "admin"
	UserRole       Role = "user"
	ContractorRole Role = "contractor"
	GuestRole      Role = "guest"
	UnknownRole    Role = "unknown"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	State     string `json:"state"`
	Role      string `json:"role"`
	Contact   string `json:"contact"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
