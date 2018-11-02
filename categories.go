package log

var (
	categories map[string] Level = make(map[string] Level)
)

func GetLevelForCategory(category string) Level {
	if level, ok := categories[category]; ok {
		return level
	} 

	return ""
}

func SetLevelForCategory(category string, level Level) {
	categories[category] = level
}

