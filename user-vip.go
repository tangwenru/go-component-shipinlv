package component_shipinlv

import (
	"fmt"
	typeChatgpt "github.com/tangwenru/go-component-shipinlv/subsystem/chatgpt/type"

	//component_shipinlv_lib "component_shipinlv/lib"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
)

type UserVip struct {
}

type UserVipDetail struct {
	Id                int64
	UserId            int64
	ProductType       string
	RatioDownPercent1 float64
	RatioDownPercent2 float64
	VipLevel          int64
	IsVip             bool
	VipDetail         typeChatgpt.VipDetail `json:"vipDetail"`
	Expired           int64
	Created           int64
	Updated           int64
}

type UserVipDetailQuery struct {
	ProductType string `json:"productType"`
}

func init() {

}

func (this *UserVip) Detail(userId int64, productType string) (error, *UserVipDetail) {
	query := UserVipDetailQuery{
		ProductType: productType,
	}
	userVipDetail := UserVipDetail{}

	bytesResult, err := component_shipinlv_lib.MainSystem(userId, "userVip/detail", &query, &userVipDetail)

	if err != nil {
		fmt.Println("UserVip) Detail:", string(bytesResult))
		//json.Unmarshal(bytesResult, &userVipDetail)
	}

	return err, &userVipDetail
}

//func (this *UserVip) List(userId int64) *global.DataResultModel {
//	bytesResult, err := component_shipinlv_lib.ApiGet(userId, "userVip/list", "")
//	userVipDetail := global.DataResultModel{}
//
//	if err != nil {
//		json.Unmarshal(bytesResult, &userVipDetail)
//	}
//
//	return &userVipDetail
//}
