package graph

import (
	"database/sql"
	"github.com/google/generative-ai-go/genai"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB    *sql.DB
	GenAi *genai.GenerativeModel
}
