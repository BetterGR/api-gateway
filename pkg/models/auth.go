package models

// TokenInfo contains the verified token information
type TokenInfo struct {
	Active    bool   `json:"active"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
}

// LoginResponse represents the response from Keycloak's token endpoint
// as specified in OAuth 2.0 RFC 6749 Section 5.1
type LoginResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	Scope            string `json:"scope"`
}

// LoginRequest represents the credentials sent for password grant login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
