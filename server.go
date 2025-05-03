package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BetterGR/api-gateway/graph"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func init() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it. Using environment variables.")
	}
}

func main() {
	// Get API Gateway configuration
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize resolver with gRPC clients
	resolver, err := graph.NewResolver()
	if err != nil {
		log.Fatalf("Failed to initialize resolver: %v", err)
	}

	// Ensure we close gRPC connections on shutdown
	defer resolver.Close()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Create an HTTP server
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: nil, // Use DefaultServeMux
	}

	// Handle graceful shutdown
	go func() {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")

	// Create a deadline to wait for current connections to complete
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Shutdown the HTTP server
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
