package lib

import (
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

func MakeSid(aesKey, roleType string, roleId, expire int64) string {
	// 999_u:123_999 , 随机数，是为了解决 AES 加密部分不变的 "不友好"
	randNum1 := RandWord(5, "")
	randNum2 := RandWord(5, "")
	//text := randNum1 + "_" + roleType + ":" + global.Int64ToString(roleId) + "_" + randNum2
	text := fmt.Sprintf("%s_%s:%d:%d_%s",
		randNum1,
		roleType,
		roleId,
		expire,
		randNum2,
	)
	return EncryptText(text, aesKey)
}
