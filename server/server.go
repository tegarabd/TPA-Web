package main

import (
	"log"
	"net/http"
	"os"
	"server/database"
	"server/directives"
	"server/graph/generated"
	"server/graph/resolver"
	"server/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initRouter() (*mux.Router) {
	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)

	c := generated.Config{Resolvers: &resolver.Resolver{}}
	c.Directives.Authenticated = directives.Authenticated
	c.Directives.Notauthenticated = directives.NotAuthenticated

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/graph"))
	router.Handle("/graph", srv)

	return router
}

const defaultPort = "8080"

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.Connect()
	database.Migrate()

	router := initRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins	:= handlers.AllowedOrigins([]string{"http://127.0.0.1:5173", "http://localhost:5173"})
	methods := handlers.AllowedMethods([]string{"POST"})

	log.Printf("connect to http://127.0.0.1:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(origins, headers, methods)(router)))

	defer database.CloseDB()
}
