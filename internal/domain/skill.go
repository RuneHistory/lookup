package domain

type Skill struct {
	Name          string
	FormattedName string
	Rank          int
	Level         int
	Experience    int
}

func NewSkill(name string, rank int, level int, experience int) *Skill {
	return &Skill{
		Name:          name,
		FormattedName: FormatSkillName(name),
		Rank:          rank,
		Level:         level,
		Experience:    experience,
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

func FormatSkillName(name string) string {
	switch name {
	case "overall":
		return "Overall"
	case "attack":
		return "Attack"
	case "defence":
		return "Defence"
	case "strength":
		return "Strength"
	case "hitpoints":
		return "Hitpoints"
	case "ranged":
		return "Ranged"
	case "prayer":
		return "Prayer"
	case "magic":
		return "Magic"
	case "cooking":
		return "Cooking"
	case "woodcutting":
		return "Woodcutting"
	case "fletching":
		return "Fletching"
	case "fishing":
		return "Fishing"
	case "firemaking":
		return "Firemaking"
	case "crafting":
		return "Crafting"
	case "smithing":
		return "Smithing"
	case "mining":
		return "Mining"
	case "herblore":
		return "Herblore"
	case "agility":
		return "Agility"
	case "theiving":
		return "Theiving"
	case "slayer":
		return "Slayer"
	case "farming":
		return "Farming"
	case "runecraft":
		return "Runecraft"
	case "hunter":
		return "Hunter"
	case "construction":
		return "Construction"
	}
	return ""
}
