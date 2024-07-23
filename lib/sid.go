package lib

import (
	"errors"
	"fmt"
)

// roleType = "u" | "s" | "c" | "w"; user, staff, client, web
type Sid struct {
	//beeController beego.Controller
}

type SidDetailResult struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Code    string    `json:"code"`
	Data    SidDetail `json:"data"`
}

type SidDetail struct {
	UserId int64 `json:"userId"`
}

func GetRoleId(roleType string, sid string) (int64, error) {
	sidDetailResult := SidDetailResult{}
	query := map[string]string{
		"roleType": roleType,
		"sid":      sid,
	}
	_, err := MainSystem(0, "user/sidInfo", &query, &sidDetailResult)

	if err != nil {
		fmt.Println("shiPinLv GetRoleId err:", err)
		return 0, err
	}

	if !sidDetailResult.Success {
		return 0, errors.New(sidDetailResult.Message)
	}

	userDetail := sidDetailResult.Data

	return userDetail.UserId, nil
}
