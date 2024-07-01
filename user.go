package app

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"` // Password minimal 6 karakter
}
