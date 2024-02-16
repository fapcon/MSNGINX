package router

import (
	cnt "geo/internal/controller"
	"github.com/go-chi/chi"
)

func Route(cnt *cnt.HandleGeo) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/address/search", cnt.SearchHandle)
	r.Post("/api/address/search", cnt.GeocodeHandle)

	return r
}
