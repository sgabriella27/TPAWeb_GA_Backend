package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sgabriella27/TPAWebGA_Back/database"
	"github.com/sgabriella27/TPAWebGA_Back/graph"
	"github.com/sgabriella27/TPAWebGA_Back/graph/generated"
	"github.com/sgabriella27/TPAWebGA_Back/graph/model"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.AllowAll().Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.Handle("/game/assets/{id}", http.HandlerFunc(ShowFile))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func ShowFile(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Print(err)
	}
	game := model.Game{ID: int64(i)}
	if err := database.GetDatabase().Preload("GameGameBanner").First(&game).Error; err != nil {
		log.Print(err)}
	writer.Write(game.GameGameBanner.ImageVideo)
}
