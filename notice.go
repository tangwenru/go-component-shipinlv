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
	UserIdList  []string `json:"userIdList"`
	PartyIdList []string `json:"partyIdList"`
	ChatId      string   `json:"chatId"`
	Text        string   `json:"text"`
}

type NoticeCreateResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Notice(query *NoticeCreateQuery) error {
	noticeToken, _ := web.AppConfig.String("AesKey.notice")
	if noticeToken == "" {
		return errors.New("请配置：AesKey.notice")
	}

	token := lib.MakeToken(noticeToken)

	req := httplib.Post("https://api-notice.shipinlv.com/notice/create?_token_=" + token)
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
