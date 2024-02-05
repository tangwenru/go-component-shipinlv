package typeChatgpt

type VipPayedUpdateQuery struct {
	ProductType  string  `json:"productType"`
	DollarAmount float64 `json:"dollarAmount"`
}

type VipPayedUpdateData struct {
}

type VipDetail struct {
	Id                int64                `json:"id"`
	DealerId          int64                `json:"dealerId"`
	Price             float64              `json:"price"`
	Month             float64              `json:"month"`
	ProductType       string               `json:"productType"`
	ProductName       string               `json:"productName"`
	MonthTitle        string               `json:"monthTitle"`
	MarketPrice       float64              `json:"marketPrice"`
	RatioDownPercent1 float64              `json:"ratioDownPercent1"`
	RatioDownPercent2 float64              `json:"ratioDownPercent2"`
	VipLevel          int64                `json:"vipLevel"`
	PropertyAmount    string               `json:"propertyAmount"`
	Recommend         int64                `json:"recommend"`
	UserPromoteInfo   UserPromoteLevelList `json:"userPromoteInfo"`
	Alt               string               `json:"alt"`
}
