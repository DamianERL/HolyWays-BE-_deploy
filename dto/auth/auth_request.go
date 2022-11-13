package authdto

type RegisterRequest struct {
	Email    string `validate:"required"  form:"email" json:"email"`
	Password string `validate:"required"  form:"password" json:"password"`
	Name     string `validate:"required"  form:"name" json:"name"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
