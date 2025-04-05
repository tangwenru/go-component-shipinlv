package subsystemLongVIdeoSales

import (
	"errors"
	"fmt"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
	typeClothLogo "github.com/tangwenru/go-component-shipinlv/subsystem/cloth-logo/type"
)

type UsageLimit struct {
}

type UsageLimitDetail struct {
	Id                int64
	UserId            int64
	ProductType       string
	RatioDownPercent1 float64
	RatioDownPercent2 float64
	UsageLimitLevel   int64
	UsageLimitId      int64
	IsUsageLimit      bool
	Expired           int64
	Created           int64
	Updated           int64
}

type UsageLimitCreateQuery struct {
}

func init() {

}

// 暂时不用了，main 直接访问
func (this *UsageLimit) Create(userId int64, query *UsageLimitCreateQuery) error {

	result := typeClothLogo.UsageLimitCreateResult{}

	bytesResult, err := component_shipinlv_lib.SubsystemClothLogo(
		userId,
		"system/usageLimit/create",
		&query,
		&result,
	)

	if err != nil {
		fmt.Println("User UsageLimit edit :", string(bytesResult))
		return err
	}

	if !result.Success {
		return errors.New(result.Message)
	}

	return nil
}
