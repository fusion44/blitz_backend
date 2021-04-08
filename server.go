package main

import (
	"log"
	"net/http"
	"time"

	loader "github.com/fusion44/raspiblitz-backend/db/loaders"
	"github.com/fusion44/raspiblitz-backend/db/repositories"
	"github.com/fusion44/raspiblitz-backend/domain"
	service "github.com/fusion44/raspiblitz-backend/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fusion44/raspiblitz-backend/db"
	"github.com/fusion44/raspiblitz-backend/graph/generated"
	"github.com/fusion44/raspiblitz-backend/graph/resolver"
	projMiddleware "github.com/fusion44/raspiblitz-backend/middleware"

	gcontext "github.com/fusion44/raspiblitz-backend/context"
)

// AppConfig holds the global configuration
var AppConfig *gcontext.Config

func main() {
	AppConfig := gcontext.LoadConfig(".")
	logger := service.NewLogger(AppConfig)

	DB := db.New(AppConfig)
	// DB.AddQueryHook(db.Logger{})
	defer DB.Close()

	userRepo := repositories.UsersRepository{DB: DB}
	infoRepo := repositories.BlitzInfoRepository{}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)
	router.Use(middleware.RequestID)
	// router.Use(middleware.Logger)
	router.Use(projMiddleware.AuthMiddleware(AppConfig, &userRepo))
	router.Use(projMiddleware.LoggerMiddleware(logger))
	router.Use(gcontext.ConfigMiddleware(AppConfig))

	c := generated.Config{Resolvers: &resolver.Resolver{
		Domain: domain.NewDomain(userRepo, infoRepo),
	}}

	queryHander := handler.New(generated.NewExecutableSchema(c))
	queryHander.AddTransport(transport.POST{})
	queryHander.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	queryHander.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", loader.UserLoaderMiddleware(DB, queryHander))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", AppConfig.ServerPort)
	log.Fatal(http.ListenAndServe(":"+AppConfig.ServerPort, router))
}
