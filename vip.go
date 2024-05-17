package component_shipinlv

import (
	"fmt"
	//component_shipinlv_lib "component_shipinlv/lib"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
)

type Vip struct {
}

type VipList struct {
	Id                    int64   `json:"id"`
	DealerId              int64   `json:"dealerId"`
	Title                 string  `json:"title"`
	Price                 float64 `json:"price"`
	Month                 float64 `json:"month"`
	ProductType           string  `json:"productType"`
	ProductTypeOrderIndex int64   `json:"productTypeOrderIndex"`
	MonthTitle            string  `json:"monthTitle"`
	MarketPrice           float64 `json:"marketPrice"`
	RatioDownPercent1     float64 `json:"ratioDownPercent1"`
	RatioDownPercent2     float64 `json:"ratioDownPercent2"`
	Alt                   string  `json:"alt"`
	VipLevel              int64   `json:"vipLevel"`
	PropertyAmount        string  `json:"propertyAmount"` // 最大值
	Recommend             int64   `json:"recommend"`
	ShowContent           string  `json:"showContent"`
	Enabled               bool    `json:"enabled"`
}

type VipDetailQuery struct {
	ProductType string `json:"productType"`
}

func init() {

}

func (this *Vip) ListByProductType(userId int64, productType string) (error, *[]VipList) {
	query := VipDetailQuery{
		ProductType: productType,
	}
	userVipDetail := make([]VipList, 0)

	bytesResult, err := component_shipinlv_lib.MainSystem(userId, "vip/listByProductType", &query, &userVipDetail)

	if err != nil {
		fmt.Println("Vip ListByProductType:", string(bytesResult))
		//json.Unmarshal(bytesResult, &userVipDetail)
	}

	return err, &userVipDetail
}
