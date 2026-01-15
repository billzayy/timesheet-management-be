package dto

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshDTO struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RespLoginDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"type"`
	ExpiredTime  int    `json:"expired_time"`
}
