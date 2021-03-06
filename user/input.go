package user

// RegisterUserInput is ...
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// LoginInput is ...
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// CheckEmailInput is ...
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
