package graph

import (
	"context"

	"github.com/BetterGR/api-gateway/graph/model"
)

// ProcessChatMessage is the resolver for the processChatMessage field.
func (r *mutationResolver) ProcessChatMessage(ctx context.Context, input model.ChatHistoryInput) (*model.ChatResponse, error) {
	// Create an authenticated context with the token
	authCtx := r.CreateAuthContext(ctx)

	// Process the chat history with ChatBotClient
	return r.ChatBotClient.ProcessMessage(authCtx, &input)
}
