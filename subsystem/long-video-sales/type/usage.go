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
}

type UsageDetail struct {
	Id               int64  `json:"id"`
	ProductType      string `json:"productType"`
	ProductName      string `json:"productName"`
	CanUseCount      int64  `json:"canUseCount"`
	AlreadyUsedCount int64  `json:"alreadyUsedCount"`
}

type UsageDetailResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    UsageDetail `json:"data"`
}
