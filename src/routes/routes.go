package routes

import (
	"my-app/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/album", controller.GetAllAlbuns).
		Methods(http.MethodGet)
	r.HandleFunc("/album", controller.CreateAlbum).
		Methods(http.MethodPost)
	r.HandleFunc("/album/{id}", controller.GetAlbum).
		Methods(http.MethodGet)
	r.HandleFunc("/album/{id}", controller.UpdateAlbum).
		Methods(http.MethodPatch)
	r.HandleFunc("/album/{id}", controller.DeleteAlbum).
		Methods(http.MethodDelete)

	return r
}
