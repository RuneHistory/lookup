package domain

type Skill struct {
	Name       string
	Rank       int
	Level      int
	Experience int
}

func NewSkill(name string, rank int, level int, experience int) *Skill {
	return &Skill{
		Name:       name,
		Rank:       rank,
		Level:      level,
		Experience: experience,
	}
}

func OrderedSkills() []string {
	return []string{
		"overall", "attack", "defence", "strength", "hitpoints",
		"ranged", "prayer", "magic", "cooking", "woodcutting",
		"fletching", "fishing", "firemaking", "crafting", "smithing",
		"mining", "herblore", "agility", "theiving", "slayer",
		"farming", "runecraft", "hunter", "construction",
	}
}
