package dto

import (
	"github.com/dev-hyunsang/clone-stackbuck-backend/models"
	"time"
)

type RequestSignUp struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	Birthday       string `json:"birthday"`
	Name           string `json:"name"`
	DisplayName    string `json:"display_name"`
	AllowMarketing bool   `json:"allow_marketing"`
}

type ResponseSignUp struct {
	Status      Status       `json:"status"`
	Data        models.Users `json:"data"`
	RespondedAt time.Time    `json:"responded_at"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	Status      Status    `json:"status"`
	Data        LoginData `json:"data"`
	RespondedAt time.Time `json:"responded_at"`
}

type LoginData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
