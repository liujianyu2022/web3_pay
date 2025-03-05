package api

type RegisterRequest struct {
	Email      string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Nickname   string `json:"nickname" example:"alan"`
	Password   string `json:"password" binding:"required" example:"123456"`
	RePassword string `json:"re_password"  example:"123456"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}