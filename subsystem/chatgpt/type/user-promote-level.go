package typeChatgpt

type UserPromoteLevelList struct {
	Title   string `json:"title"`
	Level   int64  `json:"level"`
	Color   string `json:"color"`
	Enabled bool   `json:"enabled"`
}

type UserPromoteLevelDictByLevelData map[int64]UserPromoteLevelList
