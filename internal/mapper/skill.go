package mapper

import (
	"lookup/internal/domain"
)

type SkillHttpV1 struct {
	Name       string `json:"name"`
	Rank       int    `json:"rank"`
	Level      int    `json:"level"`
	Experience int    `json:"experience"`
}

func SkillToHttpV1(s *domain.Skill) *SkillHttpV1 {
	return &SkillHttpV1{
		Name:       s.FormattedName,
		Rank:       s.Rank,
		Level:      s.Level,
		Experience: s.Experience,
	}
}
