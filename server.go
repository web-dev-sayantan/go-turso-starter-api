package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/generative-ai-go/genai"
	"github.com/ishanz23/go-turso-starter-api/db"
	"github.com/ishanz23/go-turso-starter-api/graph"
	"github.com/ishanz23/go-turso-starter-api/scripts"
	"google.golang.org/api/option"
)

const defaultPort = "8080"

func Server() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GENAI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-pro")
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db.DB, GenAi: model}}))
	err = scripts.CleanAvailability()
	if err != nil {
		log.Printf("Error cleaning availability: %v", err)
	} else {
		// Start from current date and generate for 30 days
		startDate := time.Now().Truncate(24 * time.Hour)
		endDate := startDate.AddDate(0, 0, 30)

		err = scripts.GenerateAvailability(startDate, endDate)
		if err != nil {
			log.Printf("Error generating availability: %v", err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
