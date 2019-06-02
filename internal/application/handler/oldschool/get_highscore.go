package oldschool

import (
	"fmt"
	"lookup/internal/application/service"
	"lookup/internal/domain/highscore"
	"lookup/internal/errs"
)

type GetHighScoreRequest struct {
	Nickname string
}

type GetHighScoreResponse struct {
	HighScore *highscore.HighScore
}

func NewGetHighScoreHandler(highScoreService service.HighScore) *GetHighScoreHandler {
	return &GetHighScoreHandler{
		highScoreService: highScoreService,
	}
}

type GetHighScoreHandler struct {
	highScoreService service.HighScore
}

func (h *GetHighScoreHandler) Handle(r interface{}) (interface{}, error) {
	req := r.(*GetHighScoreRequest)
	highScore, err := h.highScoreService.GetByNickname(req.Nickname)
	if err != nil {
		return nil, err
	}
	if highScore == nil {
		return nil, errs.NotFound(fmt.Sprintf("Account %s not found", req.Nickname))
	}
	return &GetHighScoreResponse{
		HighScore: highScore,
	}, nil
}
