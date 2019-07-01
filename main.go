package main

import (

	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/novacloudcz/graphql-orm/events"
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


	eventController, err := events.NewEventController()
	if err != nil {
		panic(err)
	}

	gqlHandler := handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: NewResolver(db, &eventController)}))

	playgroundHandler := handler.Playground("GraphQL playground", "/graphql")
	http.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		principalID := getPrincipalID(req)
		ctx := context.WithValue(req.Context(), gen.KeyPrincipalID, principalID)
		req = req.WithContext(ctx)
		if req.Method == "GET" {
			playgroundHandler(res, req)
		} else {
			gqlHandler(res, req)
		}
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

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPrincipalID(req *http.Request) *string {
	pID := req.Header.Get("principal-id")
	if pID == "" {
		return nil
	}
	return &pID
}

