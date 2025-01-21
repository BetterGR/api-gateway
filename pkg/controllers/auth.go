package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/BetterGR/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"k8s.io/klog/v2"
)

func extractUserInfo(tokenString string) (string, string, error) {
	parser := jwt.Parser{
		SkipClaimsValidation: true,
	}

	token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", "", fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("invalid token claims")
	}

	// Extract user ID from sub claim
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", "", fmt.Errorf("no user ID in token")
	}

	// Extract realm_access roles
	realmAccess, ok := claims["realm_access"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("no realm_access in token")
	}

	roles, ok := realmAccess["roles"].([]interface{})
	if !ok || len(roles) == 0 {
		return "", "", fmt.Errorf("no roles found in token")
	}

	// Get first role
	role, ok := roles[0].(string)
	if !ok {
		return "", "", fmt.Errorf("invalid role format")
	}

	klog.Infof("Extracted userID: %s, role: %s", userID, role)
	return role, userID, nil
}

func HandleCallback(c *gin.Context) {
	// Add CORS headers
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	// Parse request body
	var request struct {
		Code        string `json:"code"`
		RedirectURI string `json:"redirect_uri"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		klog.Errorf("Invalid request format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if request.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No code provided"})
		return
	}

	// Exchange code for token
	tokenResponse, err := exchangeCodeForToken(request.Code)
	if err != nil {
		klog.Errorf("Token exchange failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Extract user info from token
	role, userID, err := extractUserInfo(tokenResponse.AccessToken)
	if err != nil {
		klog.Errorf("Failed to extract user info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process token"})
		return
	}

	// Return the response
	response := gin.H{
		"access_token":  tokenResponse.AccessToken,
		"refresh_token": tokenResponse.RefreshToken,
		"token_type":    tokenResponse.TokenType,
		"expires_in":    tokenResponse.ExpiresIn,
		"role":          role,
		"user_id":       userID,
	}

	klog.Infof("Token exchange successful for user ID: %s with role: %s", userID, role)
	c.JSON(http.StatusOK, response)
}

// exchangeCodeForToken exchanges an authorization code for access and refresh tokens.
// It sends a request to Keycloak with the code, client credentials, and redirect URI.
// Returns the token response containing access_token, refresh_token and metadata,
// or an error if the exchange fails.
func exchangeCodeForToken(code string) (*models.LoginResponse, error) {
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURI := "http://localhost:3000/callback"

	tokenURL := fmt.Sprintf("%s/realms/betterGR/protocol/openid-connect/token", keycloakURL)

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", "account")
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		return nil, fmt.Errorf("token request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResponse models.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %v", err)
	}

	return &tokenResponse, nil
}

// verifyTokenInternal validates a token with Keycloak and returns user information
func verifyTokenInternal(token string) (*models.TokenInfo, error) {
	// Get Keycloak configuration
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	clientID := "account"
	clientSecret := os.Getenv("CLIENT_SECRET")

	// Call Keycloak's token introspection endpoint
	introspectURL := fmt.Sprintf("%s/realms/betterGR/protocol/openid-connect/token/introspect", keycloakURL)
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", introspectURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create introspection request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token introspection request failed: %v", err)
	}
	defer resp.Body.Close()

	var introspectResponse struct {
		Active   bool   `json:"active"`
		Exp      int64  `json:"exp"`
		Username string `json:"username"`
		ClientID string `json:"client_id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&introspectResponse); err != nil {
		return nil, fmt.Errorf("failed to decode introspection response: %v", err)
	}

	// Check if token is active
	if !introspectResponse.Active {
		return nil, fmt.Errorf("token is inactive")
	}

	// Verify token expiration
	if time.Now().Unix() > introspectResponse.Exp {
		return nil, fmt.Errorf("token has expired")
	}

	// Verify client ID
	if introspectResponse.ClientID != clientID {
		return nil, fmt.Errorf("invalid client ID")
	}

	// Extract additional user info from token
	role, userID, err := extractUserInfo(token)
	if err != nil {
		return nil, fmt.Errorf("failed to extract user info: %v", err)
	}

	// Return token info
	return &models.TokenInfo{
		Active:    true,
		Username:  introspectResponse.Username,
		Role:      role,
		UserID:    userID,
		ExpiresAt: introspectResponse.Exp,
	}, nil
}
