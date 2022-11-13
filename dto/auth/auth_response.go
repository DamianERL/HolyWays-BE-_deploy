package authdto

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}
