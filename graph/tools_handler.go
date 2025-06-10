package graph

import (
	"encoding/json"
	"log"
	"net/http"
)

// ToolHandler handles HTTP requests for tool operations
type ToolHandler struct {
	resolver *Resolver
}

// NewToolHandler creates a new tool handler with the given resolver
func NewToolHandler(resolver *Resolver) *ToolHandler {
	return &ToolHandler{resolver: resolver}
}

// HandleListTools handles requests to list all available tools
func (h *ToolHandler) HandleListTools(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get tools as JSON
	toolsJSON, err := h.resolver.ToolRegistry.GetToolsAsJSON()
	if err != nil {
		log.Printf("Error getting tools as JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(toolsJSON))
}

// HandleExecuteTool handles requests to execute a tool
func (h *ToolHandler) HandleExecuteTool(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req struct {
		ToolName string         `json:"tool_name"`
		Params   map[string]any `json:"params"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// Get authentication token from request context
	token := GetAuthToken(r.Context())

	// Execute the tool with the auth token
	result, err := h.resolver.ToolRegistry.ExecuteTool(req.ToolName, req.Params, token)
	if err != nil {
		log.Printf("Error executing tool %s: %v", req.ToolName, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"result":  result,
	})
}

// RegisterToolHandlers registers all tool-related HTTP handlers
func RegisterToolHandlers(resolver *Resolver, mux *http.ServeMux) {
	handler := NewToolHandler(resolver)

	// The /tools endpoint for listing available tools doesn't require auth
	mux.HandleFunc("/tools", handler.HandleListTools)

	// The /tools/execute endpoint requires authentication
	mux.Handle("/tools/execute", AuthMiddleware(http.HandlerFunc(handler.HandleExecuteTool)))
}
