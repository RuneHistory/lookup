package http_transport

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"lookup/internal/application/handler/oldschool"
	"lookup/internal/application/service"
	"net/http"
)

func WrapEndpoint(endpoint Endpoint, decoder DecoderFunc, encoder EncoderFunc, responder ResponderFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := decoder(r)
		if err != nil {
			SendError(err, w)
			return
		}
		res, err := endpoint.Handle(req)
		if err != nil {
			SendError(err, w)
			return
		}
		encoded, err := encoder(res)
		if err != nil {
			SendError(err, w)
			return
		}
		responder(w, encoded)
	})
}

func Bootstrap(r *chi.Mux, highScoreService service.HighScore) {
	getOsHighScore := oldschool.NewGetHighScoreHandler(highScoreService)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(2))
	r.Use(middleware.RequestID)
	r.Use(middleware.RedirectSlashes)

	r.Get("/osrs/highscore/{nickname}", WrapEndpoint(
		getOsHighScore,
		GetHighScoreDecoder,
		GetHighScoreEncoder,
		GetHighScoreResponder,
	))
}
