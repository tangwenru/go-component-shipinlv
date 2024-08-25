package subsystemLongVIdeoSales

import (
	"fmt"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
	typeLongVideoSales "github.com/tangwenru/go-component-shipinlv/subsystem/long-video-sales/type"
)

type Usage struct {
}

type UsageDetail struct {
	Id                int64
	UserId            int64
	ProductType       string
	RatioDownPercent1 float64
	RatioDownPercent2 float64
	UsageLevel        int64
	UsageId           int64
	IsUsage           bool
	Expired           int64
	Created           int64
	Updated           int64
}

func init() {

}

// @timeType: 'day' | 'month'
func (this *Usage) List(userId int64, productType, timeType string) (error, *[]typeLongVideoSales.UsageList) {
	query := typeLongVideoSales.UsageListQuery{
		ProductType: productType,
		TimeType:    timeType,
	}
	result := typeLongVideoSales.UsageListResult{}

	bytesResult, err := component_shipinlv_lib.SubsystemLongVideoSales(
		userId,
		"system/product/usage",
		&query,
		&result,
	)

	if err != nil {
		fmt.Println("UserUsage Detail :", string(bytesResult))
		//json.Unmarshal(bytesResult, &userUsageDetail)
		return err, nil
	}

	return nil, &result.Data
}
