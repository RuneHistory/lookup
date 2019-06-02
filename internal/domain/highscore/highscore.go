package highscore

import "time"

type HighScore struct {
	Nickname  string
	CreatedAt time.Time
}

func NewHighScore(nickname string, createdAt time.Time) *HighScore {
	return &HighScore{
		Nickname:  nickname,
		CreatedAt: createdAt,
	}
}
