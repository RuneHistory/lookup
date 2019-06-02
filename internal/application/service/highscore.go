package service

import "lookup/internal/domain/highscore"

type HighScore interface {
	GetByNickname(nickname string) (*highscore.HighScore, error)
}

func NewHighScoreService() HighScore {
	return &HighScoreService{}
}

type HighScoreService struct {
}

func (s *HighScoreService) GetByNickname(id string) (*highscore.HighScore, error) {
	return nil, nil
}
