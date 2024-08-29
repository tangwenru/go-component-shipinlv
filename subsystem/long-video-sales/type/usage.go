package typeLongVideoSales

type UsageListQuery struct {
	ProductType string `json:"productType"`
	TimeType    string `json:"timeType"`
}

type UsageList struct {
	Time        string `json:"time"`
	FailCount   int    `json:"failCount"`
	FinishCount int    `json:"finishCount"`
}

type UsageListResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    []UsageList `json:"data"`
}

///////

type UsageDetailQuery struct {
	ProductType string `json:"productType"`
	UserId      int64  `json:"userId"`
}

type UsageDetail struct {
	UserId                  int64  `json:"userId"`
	UserIdKey               string `json:"userIdKey"`
	ProductType             string `json:"productType"`
	ProductName             string `json:"productName"`
	CanUseCount             int64  `json:"canUseCount"`
	AlreadyUsedCount        int64  `json:"alreadyUsedCount"`
	HistoryAlreadyUsedCount int64  `json:"historyAlreadyUsedCount"`
}

type UsageDetailResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    UsageDetail `json:"data"`
}

type UsageEditQuery struct {
	UserId      int64  `json:"userId"`
	ProductType string `json:"productType"`
	CanUseCount int    `json:"canUseCount"`
}

type UsageEditResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
