package diffServer

import (
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Status      string `json:"status"`
	IssuedAt    int64  `json:"issued_at"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (e *Token) SetAccessToken(accessToken string) {
	e.AccessToken = accessToken
}

func (e *Token) SetTokenType(tokenType string) {
	e.TokenType = tokenType
}

// SetStatus define o estado mostrado ao usuário na tela
func (e *Token) SetStatus(status string) {
	e.Status = status
}

func (e *Token) SetIssuedAt(issuedAt int64) {
	e.IssuedAt = issuedAt
}

func (e *Token) SetExpiresIn(expiresIn int64) {
	e.ExpiresIn = expiresIn
}

func (e *Token) GetAccessToken() string {
	return e.AccessToken
}

func (e *Token) GetTokenType() string {
	return e.TokenType
}

func (e *Token) GetStatus() string {
	return e.Status
}

func (e *Token) GetIssuedAt() int64 {
	return e.IssuedAt
}

func (e *Token) GetExpiresIn() int64 {
	return e.ExpiresIn
}

func (e *Token) Init() {
	e.AccessToken = gofakeit.UUID()
	e.TokenType = "Bearer"
	e.Status = "approved"
	e.IssuedAt = time.Now().UTC().Unix()
	e.ExpiresIn = 3599
}
