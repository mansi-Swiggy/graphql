package main

import (
	"example/api"
	"example/graph"
	"example/logging"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	r := gin.Default()
	r.GET(api.HealthCheckString, api.NewHealthCheckHandler().HealthCheck)

	zap.L().Info("connect to GraphQL playground", logging.AutoField("port", port))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		zap.L().Error("Server error", zap.Error(err))

	}
}
