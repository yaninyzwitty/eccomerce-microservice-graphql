package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/yaninyzwitty/new-galgrn-go/database"
	"github.com/yaninyzwitty/new-galgrn-go/graph"
	"github.com/yaninyzwitty/new-galgrn-go/pkg"
)

var (
	username string
	password string
	hostname string
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var cfg pkg.Config
	file, err := os.Open("config.yaml")
	if err != nil {
		slog.Error("failed to open config yaml")
		os.Exit(1)
	}

	if err := cfg.LoadConfig(file); err != nil {
		slog.Error("failed to load config")
		os.Exit(1)
	}

	err = godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env file")
	}

	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	hostname = os.Getenv("HOSTNAME")
	if username == "" || password == "" || hostname == "" {
		slog.Error("failed to load mongo db config")
		os.Exit(1)
	}

	mongoDBConfig := &database.MongoDBConfig{
		Username: username,
		Password: password,
		Hostname: hostname,
	}

	mongoDBClient, err := database.NewMongoDbConnection(ctx, mongoDBConfig)
	if err != nil {
		slog.Error("failed to create mongo db connection: ", "error", err)
		os.Exit(1)
	}
	defer mongoDBClient.Disconnect(ctx)

	err = database.PingDatabase(ctx, mongoDBClient)
	if err != nil {
		slog.Error("failed to ping database: ", "error", err)
		os.Exit(1)
	}

	slog.Info("connected to mongo succesfully 🚀")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		MongoDBClient: mongoDBClient,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	stopCH := make(chan os.Signal, 1)
	signal.Notify(stopCH, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("server is listening on :" + fmt.Sprintf("%d", cfg.Server.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to listen and serve", "error", err)
		}
	}()
	<-stopCH

	slog.Info("shuttting down the server...")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shut down the server", "error", err)
	} else {
		slog.Info("server stopped down gracefully")
	}

}
