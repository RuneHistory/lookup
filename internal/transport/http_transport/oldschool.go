package http_transport

import (
	"github.com/go-chi/chi"
	"github.com/jmwri/go-http"
	"lookup/internal/application/handler/oldschool"
	"lookup/internal/mapper"
	"net/http"
)

func GetHighScoreDecoder(r *http.Request) (interface{}, error) {
	req := &oldschool.GetHighScoreRequest{
		Nickname: chi.URLParam(r, "nickname"),
	}
	return req, nil
}

func GetHighScoreEncoder(d interface{}) (interface{}, error) {
	res := d.(*oldschool.GetHighScoreResponse)
	return mapper.HighScoreToHttpV1(res.HighScore), nil
}

func GetHighScoreResponder(w http.ResponseWriter, d interface{}) {
	go_http.SendJson(d, w)
}
