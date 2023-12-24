package generic_handlers

import (
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data/dao"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Make props private
type GenericHandler[T types.DocumentModel] struct {
	DAO             *dao.DAO[T]
	Router          chi.Router
	Validate        *validator.Validate
	Collection      string
	FilterableProps *[]string
}

func New[T types.DocumentModel](db *mongo.Database, router chi.Router, validate *validator.Validate, collection string, filterableProps *[]string) *GenericHandler[T] {
	return &GenericHandler[T]{
		DAO:             dao.New[T](db),
		Router:          router,
		Validate:        validate,
		Collection:      collection,
		FilterableProps: filterableProps,
	}
}

func (g *GenericHandler[T]) ImplementGenericCrud(routeSegment string) {
	basePath := path.Join("/", routeSegment)
	pathWithIdQuery := path.Join(basePath, "/{id}")

	g.Router.Get(basePath, g.GetListHandler)

	g.Router.Get(pathWithIdQuery, g.GetByIDHandler)

	g.Router.Post(basePath, g.PostDocumentHandler)

	g.Router.Delete(pathWithIdQuery, g.DeleteDocumentByIDHandler)

	g.Router.Put(pathWithIdQuery, g.UpdateDocumentHandler)
}
