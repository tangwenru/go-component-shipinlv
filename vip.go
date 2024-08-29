package component_shipinlv

import (
	"errors"
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
	CanBuy                bool    `json:"canBuy"`
	Enabled               bool    `json:"enabled"`
}

type VipDetailQuery struct {
	ProductType string `json:"productType"`
	Domain      string `json:"domain"`
}

type VipListResult struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []VipList `json:"data"`
}

func init() {

}

func (this *Vip) ListByProductType(userId int64, productType string, domain string) (error, *[]VipList) {
	query := VipDetailQuery{
		ProductType: productType,
		Domain:      domain,
	}
	vipListResult := VipListResult{}

	bytesResult, err := component_shipinlv_lib.MainSystem(userId, "vip/listByProductType", &query, &vipListResult)

	if err != nil {
		fmt.Println("Vip List By Product Type err:", string(bytesResult), err)
	}

	if !vipListResult.Success {
		return errors.New(vipListResult.Message), nil
	}

	return err, &vipListResult.Data
}
