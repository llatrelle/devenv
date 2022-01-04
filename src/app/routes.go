package app

import (
	"devEnv/src/api"
	"devEnv/src/api/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.NoCache)

	r.Mount("/api", func() chi.Router {
		r := chi.NewRouter()
		r.Get("/environments", controller.GetEnvironments)
		r.Post("/environments", controller.GetEnvironments)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(api.IDCtx)
			r.Patch("/", controller.GetEnvironments)  // PUT /articles/123
			r.Delete("/", controller.GetEnvironments) // DELETE /articles/123
		})

		return r

	}())
	return r
}
