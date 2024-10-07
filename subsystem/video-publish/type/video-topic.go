package typeVideoTopic

type VideoTopicTopicList struct {
	Id      int64   `json:"id"`
	GroupId int64   `json:"groupId"`
	Title   float64 `json:"title"`
	Enabled bool    `json:"enabled"`
	Created int64   `json:"created"`
}

type VideoTopicTopicListQuery struct {
	GroupId int64 `json:"groupId,omitempty"`
	Enabled bool  `json:"enabled,omitempty"`
}

type VideoTopicListResult struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Data    []VideoTopicTopicList `json:"data"`
}
