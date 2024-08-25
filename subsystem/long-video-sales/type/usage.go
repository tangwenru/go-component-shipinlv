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
