package handlers

import (
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Finish this setup in order to use Dependency Injection

type Handler interface {
	RegisterRoutes(r *chi.Router)
}

type RouteHandler struct {
	Controllers []Handler
	DBConnect   *mongo.Database
}

func (r *RouteHandler) RegisterControllerRoutes(router *chi.Router) {
	for _, x := range r.Controllers {
		x.RegisterRoutes(router)
	}
}
