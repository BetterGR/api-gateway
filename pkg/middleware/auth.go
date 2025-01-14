package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

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

	// Debug log the claims
	klog.Infof("Token claims: %+v", claims)

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

// LoginHandler is the handler for the login route.
func LoginHandler(c *gin.Context) {
	// Add CORS headers
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	var credentials models.LoginRequest
	klog.Info("LoginHandler started")

	clientSecret := os.Getenv("CLIENT_SECRET")
	keycloakURL := os.Getenv("KEYCLOAK_URL")

	// Debug environment variables
	klog.Infof("Keycloak URL: %s", keycloakURL)
	klog.Infof("Client Secret length: %d", len(clientSecret))

	if err := c.ShouldBindJSON(&credentials); err != nil {
		klog.Errorf("Invalid JSON input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Log received credentials (password length only for security)
	klog.Infof("Received credentials - Username: %s, Password length: %d",
		credentials.Username, len(credentials.Password))

	tokenURL := fmt.Sprintf("%s/realms/betterGR/protocol/openid-connect/token", keycloakURL)

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "api-gateway")
	data.Set("client_secret", clientSecret)
	data.Set("username", credentials.Username)
	data.Set("password", credentials.Password)
	data.Set("scope", "openid")

	// Log request details
	klog.Infof("Sending request to: %s", tokenURL)
	klog.Infof("Request data: client_id=%s, username=%s, grant_type=password, scope=openid",
		"api-gateway", credentials.Username)

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		klog.Errorf("HTTP request failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Authentication request failed: %v", err)})
		return
	}
	defer resp.Body.Close()

	// Log response status
	klog.Infof("Keycloak response status: %d", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		klog.Errorf("Failed to read response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read authentication response"})
		return
	}

	// Log response body length and first few characters
	klog.Infof("Response body length: %d", len(string(body)))
	if len(string(body)) > 50 {
		klog.Infof("Response body preview: %s...", string(body)[:50])
	}

	if resp.StatusCode != http.StatusOK {
		var keycloakError struct {
			Error            string `json:"error"`
			ErrorDescription string `json:"error_description"`
		}
		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&keycloakError); err != nil {
			c.JSON(resp.StatusCode, gin.H{"error": "Authentication failed"})
			return
		}
		c.JSON(resp.StatusCode, gin.H{"error": keycloakError.ErrorDescription})
		return
	}

	var tokenResponse models.TokenResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&tokenResponse); err != nil {
		klog.Errorf("Failed to decode token response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process response"})
		return
	}

	// Extract role and user ID from the token
	role, userID, err := extractUserInfo(tokenResponse.AccessToken)
	if err != nil {
		klog.Errorf("Failed to extract user info from token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process token"})
		return
	}

	// Return the response in the exact format the frontend expects
	response := gin.H{
		"access_token":  tokenResponse.AccessToken,
		"refresh_token": tokenResponse.RefreshToken,
		"token_type":    tokenResponse.TokenType,
		"expires_in":    tokenResponse.ExpiresIn,
		"role":          role,
		"user_id":       userID,
		"username":      credentials.Username,
	}

	klog.Infof("Sending successful response for user: %s with role: %s and ID: %s", credentials.Username, role, userID)
	c.JSON(http.StatusOK, response)
}
