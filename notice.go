package component_shipinlv

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
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

func Notice(noticeToken string, query *NoticeCreateQuery) error {
	req := httplib.Post("https://api-notice.shipinlv.com/notice/create")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	req.Header("Content-Type", "application/json")

	req.Param("_token_", noticeToken)

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
