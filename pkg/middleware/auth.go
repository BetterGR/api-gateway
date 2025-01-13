package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BetterGR/api-gateway/pkg/models"
	"github.com/BetterGR/api-gateway/pkg/utils"
	"io"
	"k8s.io/klog/v2"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// LoginHandler is the handler for the login route.
func LoginHandler(c *gin.Context) {
	var credentials models.LoginRequest
	klog.Info("LoginHandler")
	clientSecret := os.Getenv("CLIENT_SECRET")
	utils.Debug("Client secret: %s", clientSecret)

	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.Debug("Invalid JSON input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Keycloak configuration
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	tokenURL := fmt.Sprintf("%s/realms/betterGR/protocol/openid-connect/token", keycloakURL)

	utils.Debug("Attempting login for user: %s", credentials.Username)
	utils.Debug("Token URL: %s", tokenURL)

	// Prepare token request
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "api-gateway")
	data.Set("client_secret", clientSecret)
	data.Set("username", credentials.Username)
	data.Set("password", credentials.Password)
	klog.Info(credentials.Username)
	klog.Info(credentials.Password)
	utils.Debug("Sending request with data: %+v", data)

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		utils.Debug("HTTP request failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
		return
	}
	defer resp.Body.Close()

	utils.Debug("Keycloak response status: %d", resp.StatusCode)

	// Read and log the response body for debugging
	body, _ := io.ReadAll(resp.Body)
	utils.Debug("Response body: %s", string(body))

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
		utils.Debug("Failed to decode token response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process response"})
		return
	}

	utils.Debug("Login successful for user: %s", credentials.Username)
	c.JSON(http.StatusOK, tokenResponse)
}
