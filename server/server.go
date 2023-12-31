package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/unexpectedtokens/api-tester/server/db"
	"github.com/unexpectedtokens/api-tester/server/handlers"
)

func RunServer() {
	godotenv.Load()

	logger := httplog.NewLogger("quiktest-api", httplog.Options{
		Concise: false,
	})

	validate := validator.New(validator.WithRequiredStructEnabled())

	dbConnectionString := os.Getenv("MONGO_CONNECTIONSTRING")
	connection, err := db.NewConnection(dbConnectionString)

	if err != nil {
		panic(err)
	}

	defer connection.Disconnect(context.Background())

	db := db.GetDB(connection)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Register middleware
	r.Use(httplog.RequestLogger(logger, []string{"/ping"}))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer, middleware.RequestID)

	register := handlers.NewRegister(db, r, validate)
	register.RegisterHandlers()

	logger.Info().Msg("Registered the following routes")
	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	})

	logger.Info().Msg("Running on 8080")
	http.ListenAndServe("localhost:8080", r)
}
