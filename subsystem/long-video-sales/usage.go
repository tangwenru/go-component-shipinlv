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
		"system/productUsage/list",
		&query,
		&result,
	)

	if err != nil {
		fmt.Println("User Usage Detail :", string(bytesResult))
		return err, nil
	}

	return nil, &result.Data
}

func (this *Usage) Detail(userId int64, productType string) (error, *typeLongVideoSales.UsageDetail) {
	query := typeLongVideoSales.UsageDetailQuery{
		ProductType: productType,
	}
	result := typeLongVideoSales.UsageDetailResult{}

	bytesResult, err := component_shipinlv_lib.SubsystemLongVideoSales(
		userId,
		"system/productUsage/detail",
		&query,
		&result,
	)

	if err != nil {
		fmt.Println("User Usage Detail :", string(bytesResult))
		return err, nil
	}

	return nil, &result.Data
}

// 经销商 修改
func (this *Usage) Edit(userId int64, productType string, canUseCount int) error {
	query := typeLongVideoSales.UsageEditQuery{
		ProductType: productType,
		CanUseCount: canUseCount,
	}
	result := typeLongVideoSales.UsageEditResult{}

	bytesResult, err := component_shipinlv_lib.SubsystemLongVideoSales(
		userId,
		"system/productUsage/edit",
		&query,
		&result,
	)

	if err != nil {
		fmt.Println("User Usage edit :", string(bytesResult))
		return err
	}

	return nil
}
