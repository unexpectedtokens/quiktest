package router

import (
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func ImplementGenericCrud[T types.FilterableDocumentModel](r chi.Router, db *mongo.Database, validate *validator.Validate, routeSegment string, model T) {
	basePath := path.Join("/", routeSegment)
	pathWithIdQuery := path.Join(basePath, "/{id}")

	r.Get(basePath, handlers.GetListHandler[T](db, model))

	r.Get(pathWithIdQuery, handlers.GetByIDHandler[T](db, model))

	r.Post(basePath, handlers.PostDocumentHandler[T](db, validate))

	r.Delete(pathWithIdQuery, handlers.DeleteDocumentByIDHandler[T](db, model))

	r.Put(pathWithIdQuery, handlers.UpdateDocumentHandler[T](db, validate))
}
