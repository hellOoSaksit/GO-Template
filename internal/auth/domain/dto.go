package domain

//Model Login
type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

//Model Register
type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,regexp=^(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).+$"`
	Role     string `json:"role"`
}

//Model Response Login
type LoginResponse struct {
	Token string `json:"token"`
}

//Model Response Register
type RegisterResponse struct {
	Message string `json:"message"`
}

//Model Request Register
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

//Model Request Login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
