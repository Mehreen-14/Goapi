package handlers

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware" 
	"github.com/Mehreen-14/Goapi/internal/middleware"
	"net/http"
)

func Handler(r *chi.Mux) {
	//r.Get("/balance", getCoinBalance)
	r.Use(chimiddle.StripSlashes)
	
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", getCoinBalance)
		
	})
}
