package domains

type Auth struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (a *AuthRequest) TableName() string {
	return "account"
}
