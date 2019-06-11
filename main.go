package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/novacloudcz/graphql-orm-example/gen"
)

const (
	defaultPort = "80"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	urlString := os.Getenv("DATABASE_URL")
	if urlString == "" {
		panic(fmt.Errorf("missing DATABASE_URL environment variable"))
	}

	db := gen.NewDBWithString(urlString)
	defer db.Close()
	db.AutoMigrate()

	gqlHandler := handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: &gen.Resolver{DB: db}}))
	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), gen.DBContextKey, db)
		req = req.WithContext(ctx)
		gqlHandler(res, req)
	})

	http.HandleFunc("/healthcheck", func(res http.ResponseWriter, req *http.Request) {
		if err := db.Ping(); err != nil {
			res.WriteHeader(400)
			res.Write([]byte("ERROR"))
			return
		}
		res.WriteHeader(200)
		res.Write([]byte("OK"))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
