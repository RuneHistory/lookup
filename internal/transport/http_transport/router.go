package http_transport

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmwri/go-http"
	"lookup/internal/application/handler/oldschool"
	"lookup/internal/application/service"
)

func Bootstrap(r *chi.Mux, highScoreService service.HighScore) {
	getOsHighScore := oldschool.NewGetHighScoreHandler(highScoreService)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(2))
	r.Use(middleware.RequestID)
	r.Use(middleware.RedirectSlashes)

	r.Get("/osrs/highscore/{nickname}", go_http.WrapEndpoint(
		getOsHighScore,
		GetHighScoreDecoder,
		GetHighScoreEncoder,
		GetHighScoreResponder,
	))
}
