package component_shipinlv

import (
	"errors"
	"fmt"
	//component_shipinlv_lib "component_shipinlv/lib"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
)

type User struct {
}

type UserDetailResult struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Code    string     `json:"code"`
	Data    UserDetail `json:"data"`
}

type UserDetail struct {
	Id          int64  `json:"id"`
	AccountName string `json:"accountName"`
	Nickname    string `json:"nickname"`
	AvatarUrl   string `json:"avatarUrl"`
	Enabled     bool   `json:"enabled"`
	Expired     int64  `json:"expired"`
	WebSid      string `json:"webSid"`
	Role        string `json:"role"`
	AppSid      string `json:"appSid"`
	ClientSid   string `json:"clientSid"`
	DealerId    int64  `json:"dealerId"`
	Created     int64  `json:"created"`
}

func (this *User) Detail(userId int64) (error, *UserDetail) {
	userDetailResult := UserDetailResult{}
	query := map[string]string{}
	_, err := component_shipinlv_lib.MainSystem(userId, "user/info", &query, &userDetailResult)

	userDetail := UserDetail{}
	if err != nil {
		fmt.Println("shiPinLv user info err:", err)
		return err, &userDetail
	}

	if !userDetailResult.Success {
		return errors.New(userDetailResult.Message), &userDetail
	}

	return err, &userDetailResult.Data
}

// 一个用户只能登录一个客户端
func (this *User) DetailByOneClient(userId int64, productType, clientUniqueKey string) (error, *UserDetail) {
	userDetailResult := UserDetailResult{}
	query := map[string]string{
		"productType":     productType,
		"clientUniqueKey": clientUniqueKey,
	}
	_, err := component_shipinlv_lib.MainSystem(userId, "user/detailByOneClient", &query, &userDetailResult)

	userDetail := UserDetail{}
	if err != nil {
		fmt.Println("shi Pin Lv user info err:", err)
		return err, &userDetail
	}

	if !userDetailResult.Success {
		return errors.New(userDetailResult.Message), &userDetail
	}

	return err, &userDetailResult.Data
}
