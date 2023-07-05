package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irshadpalayadan/job-service-pg-based/graph/generated"
	"github.com/irshadpalayadan/job-service-pg-based/infra"
	"github.com/irshadpalayadan/job-service-pg-based/repository/postgres"
	"github.com/irshadpalayadan/job-service-pg-based/resolvers"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	writeDB, _, _ := infra.NewPostgresDB(false)
	repository := postgres.InitDBRepository(writeDB)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		WriteDB: repository,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
