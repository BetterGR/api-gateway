package graph

import (
	"context"
	"net/http"
	"strings"
)

// Key for storing auth token in context
type contextKey string

const (
	// AuthTokenKey is the key used to store the auth token in the context
	AuthTokenKey = contextKey("auth_token")
)

// AuthMiddleware extracts the JWT token from the Authorization header and adds it to the context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			// Usually the format is "Bearer token" - split to get the token part
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 {
				token := parts[1]
				// Store the token in context for resolvers to use
				ctx := context.WithValue(r.Context(), AuthTokenKey, token)
				r = r.WithContext(ctx)
			}
		}

		// Call the next handler with the updated context
		next.ServeHTTP(w, r)
	})
}

// GetAuthToken gets the auth token from the GraphQL context
func GetAuthToken(ctx context.Context) string {
	if token, ok := ctx.Value(AuthTokenKey).(string); ok {
		return token
	}
	return ""
}
