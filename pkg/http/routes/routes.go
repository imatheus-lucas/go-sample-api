package routes

import (
	"github.com/gorilla/mux"
	"github.com/imatheus-lucas/go-sample-api/pkg/http/api"
	"github.com/imatheus-lucas/go-sample-api/pkg/utils"
)

func Register(router *mux.Router) {

	routes := utils.Routes{
		utils.AddRoute("/books", "GET", api.GetBooks),
		utils.AddRoute("/book/{id}", "DELETE", api.DeleteBook),
		utils.AddRoute("/book/{id}", "PUT", api.UpdateBook),
		utils.AddRoute("/book/{id}", "GET", api.GetBook),
		utils.AddRoute("/book", "POST", api.CreateBook),
	}

	for _, r := range routes {
		router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}
}
