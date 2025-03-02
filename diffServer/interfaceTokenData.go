package diffServer

type TokenData interface {
	SetAccessToken(accessToken string)
	SetTokenType(tokenType string)

	// SetStatus define o estado mostrado ao usuário na tela
	SetStatus(status string)
	SetIssuedAt(issuedAt int64)
	SetExpiresIn(expiresIn int64)

	GetAccessToken() string
	GetTokenType() string
	GetStatus() string
	GetIssuedAt() int64
	GetExpiresIn() int64
}
