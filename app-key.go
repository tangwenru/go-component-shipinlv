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
	KeyType   string `json:"keyType"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

func init() {

}

func (this *AppKey) DetailByUserId(userId int64, keyType string) (*AppKeyDetail, error) {
	appKeyDetailResult := AppKeyDetailResult{}
	query := map[string]string{
		"keyType": keyType,
	}
	_, err := component_shipinlv_lib.MainSystem(userId, "appKey/detail", &query, &appKeyDetailResult)

	appKeyDetail := AppKeyDetail{}
	if err != nil {
		fmt.Println("app Key info err:", err)
		return &appKeyDetail, err
	}

	if !appKeyDetailResult.Success {
		return &appKeyDetail, errors.New(appKeyDetailResult.Message)
	}

	return &appKeyDetailResult.Data, err
}
