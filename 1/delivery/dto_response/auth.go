package dto_response

import (
	"myapp/model"
	"time"
)

type AuthTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiredAt time.Time `json:"access_token_expired_at"`
	TokenType            string    `json:"token_type"`
} // @name AuthTokenResponse

func NewAuthTokenResponse(token model.Token) AuthTokenResponse {
	return AuthTokenResponse{
		AccessToken:          token.AccessToken,
		AccessTokenExpiredAt: token.ExpiredAt,
		TokenType:            token.TokenType,
	}
}

func NewAuthTokenResponseP(token model.Token) *AuthTokenResponse {
	r := NewAuthTokenResponse(token)
	return &r
}
