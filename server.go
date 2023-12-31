package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irshadpalayadan/job-service-pg-based/bootstrap"
	"github.com/irshadpalayadan/job-service-pg-based/graph/generated"
	"github.com/irshadpalayadan/job-service-pg-based/infra"
	"github.com/irshadpalayadan/job-service-pg-based/repository/postgres"
	"github.com/irshadpalayadan/job-service-pg-based/resolvers"
	"github.com/irshadpalayadan/job-service-pg-based/utils"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger, _ := infra.InitLogger()
	writeDB, _, _ := infra.InitPostgresDB(false, logger)
	repository := postgres.InitDBRepository(writeDB, logger)

	bootstrap.InitBootstrap(writeDB, logger)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		WriteDB: repository,
		Logger:  logger,
	}}))
	srv.SetErrorPresenter(utils.GqlErrorRespMiddleware(logger))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
