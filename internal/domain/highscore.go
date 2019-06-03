package domain

import "time"

type HighScore struct {
	Nickname  string
	CreatedAt time.Time
	Skills    []*Skill
}

func NewHighScore(nickname string, createdAt time.Time) *HighScore {
	return &HighScore{
		Nickname:  nickname,
		CreatedAt: createdAt,
		Skills:    make([]*Skill, 24),
	}
}
