package mapper

import (
	"lookup/internal/domain"
	"time"
)

type HighScoreHttpV1 struct {
	Nickname  string                  `json:"nickname"`
	CreatedAt time.Time               `json:"created_at"`
	Skills    map[string]*SkillHttpV1 `json:"skills"`
}

func HighScoreToHttpV1(hs *domain.HighScore) *HighScoreHttpV1 {
	mapped := &HighScoreHttpV1{
		Nickname:  hs.Nickname,
		CreatedAt: hs.CreatedAt,
		Skills:    make(map[string]*SkillHttpV1),
	}

	for _, skill := range hs.Skills {
		mapped.Skills[skill.Name] = SkillToHttpV1(skill)
	}

	return mapped
}
