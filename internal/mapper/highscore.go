package mapper

import (
	"lookup/internal/domain/highscore"
	"time"
)

type HighScoreHttpV1 struct {
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

func HighScoreToHttpV1(hs *highscore.HighScore) *HighScoreHttpV1 {
	return &HighScoreHttpV1{
		Nickname:  hs.Nickname,
		CreatedAt: hs.CreatedAt,
	}
}
