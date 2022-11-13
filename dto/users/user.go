package usersdto

type UpdateUser struct{
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Image    string `json:"image" form:"image"`
}