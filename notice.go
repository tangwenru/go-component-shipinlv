package component_shipinlv

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/server/web"
	"github.com/tangwenru/go-component-shipinlv/lib"
)

type NoticeCreateQuery struct {
	PlanId int64  `json:"planId"`
	Text   string `json:"text"`
}

type NoticeCreateResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Notice(userId int64, query *NoticeCreateQuery, expire int64) error {
	noticeAes, _ := web.AppConfig.String("AesKey.notice")
	if noticeAes == "" {
		return errors.New("请配置：AesKey.notice")
	}
	runMode, _ := web.AppConfig.String("RunMode")

	apiUrl := "https://api-notice.shipinlv.com/notice/create"
	if runMode == "dev" {
		apiUrl = "http://127.0.0.1:51817"
	}

	token := lib.MakeSid(noticeAes, "u", userId, expire)

	req := httplib.Post(apiUrl + "?_token_=" + token)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	req.Header("Content-Type", "application/json")

	req.Param("_token_", token)

	dataByte, err := json.Marshal(query)
	req.Body(dataByte)

	byteResult, err := req.Bytes()
	if err != nil {
		fmt.Println("get api:", err)
		return errors.New("发送消息失败：" + err.Error())
	}

	resultData := NoticeCreateResult{}
	json.Unmarshal(byteResult, &resultData)

	if !resultData.Success {
		return errors.New("发送结果失败：" + resultData.Message)
	}

	return nil
}
