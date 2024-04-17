package component_shipinlv

import (
	"errors"
	"fmt"
	//component_shipinlv_lib "component_shipinlv/lib"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
)

type AppKey struct {
}

type AppKeyDetailResult struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Code    string       `json:"code"`
	Data    AppKeyDetail `json:"data"`
}

type AppKeyDetail struct {
	//Id int64 `json:"id"`
	//DealerId  int64 `json:"dealer_id"`
	KeyType   string `json:"keyType"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	//Enabled   bool   `json:"enabled"`
	//Created   int64  `json:"created"`
}

func init() {

}

func (this *AppKey) DetailByUserId(userId int64) (*AppKeyDetail, error) {
	appKeyDetailResult := AppKeyDetailResult{}
	query := map[string]string{}
	_, err := component_shipinlv_lib.MainSystem(userId, "appKey/detail", &query, &appKeyDetailResult)

	appKeyDetail := AppKeyDetail{}
	if err != nil {
		fmt.Println("appKey info err:", err)
		return &appKeyDetail, err
	}

	if !appKeyDetailResult.Success {
		return &appKeyDetail, errors.New(appKeyDetailResult.Message)
	}

	return &appKeyDetailResult.Data, err
}
