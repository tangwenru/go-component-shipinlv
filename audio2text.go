package component_shipinlv

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/tangwenru/go-component-shipinlv/lib"
	"strings"
)

type Audio2Text struct {
}

type Audio2TextDetail struct {
	Id      int64 `json:"id"`
	SrtText string
}

type Audio2TextDetailResult struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    Audio2TextDetail `json:"data"`
}

type Audio2TextCreateQuery struct {
	TaskId        int64  `json:"taskId"`
	TaskType      string `json:"taskType"`
	LineMaxLetter int    `json:"lineMaxLetter"`
	AudioFilePath string `json:"audioFilePath"`
}

type Audio2TextApiResult struct {
	Message string `json:"message,omitempty"`
	Data    struct {
		Query struct {
			AudioFilePath string `json:"audioFilePath"`
		} `json:"query"`
		Result []struct {
			Key       string  `json:"key"`
			Text      string  `json:"text"`
			Timestamp [][]int `json:"timestamp"`
		} `json:"result"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (this *Audio2Text) Create(userId int64, query *Audio2TextCreateQuery) (*Audio2TextDetail, error) {
	// 先找一个 服务器
	workServer := WorkServer{}
	workServerDetail, errWorkServerDetail := workServer.RandDetail(userId, "audio2text")
	if errWorkServerDetail != nil {
		return nil, errWorkServerDetail
	}

	port := fmt.Sprintf("%d", workServerDetail.Port)
	if len(port) > 1 {
		port = ":" + port
	}
	apiUrl := fmt.Sprintf("%s://%s%s/api-audio2text/audio2text",
		workServerDetail.DomainProtocol,
		workServerDetail.Domain,
		port,
	)

	req := httplib.Post(apiUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	data := map[string]interface{}{
		"audioFilePath": query.AudioFilePath,
	}

	req.Header("Content-Type", "application/json")

	dataByte, err := json.Marshal(&data)
	req.Body(dataByte)

	byteResult, err := req.Bytes()
	if err != nil {
		fmt.Println("get api :", err)
		return nil, err
	}

	resultData := Audio2TextApiResult{}
	json.Unmarshal(byteResult, &resultData)

	if !resultData.Success {
		return nil, errors.New("识别字幕失败：" + resultData.Message)
	}

	fmt.Println("api:", resultData)
	outDetail := Audio2TextDetail{
		SrtText: this.ApiResult2Srt(query.LineMaxLetter, &resultData),
	}

	return &outDetail, nil
}

func (this *Audio2Text) ApiResult2Srt(lineMaxLetter int, data *Audio2TextApiResult) string {
	srtList := make([]string, 0)
	//每行最多多少字
	if lineMaxLetter < 1 {
		lineMaxLetter = 12
	}
	for _, items := range data.Data.Result {
		textList := strings.Split(items.Text, " ")
		timestampList := items.Timestamp
		textListLen := len(textList)
		timestampListLen := len(timestampList)
		//对齐
		//文本太长
		if textListLen > timestampListLen {
			textList = textList[0:timestampListLen]
		} else if textListLen < timestampListLen {
			timestampList = timestampList[0:textListLen]
		}
		//更新
		textListLen = len(textList)
		timestampListLen = len(timestampList)

		lineBreakMillisecond := 500
		lineText := make([]string, 0)
		lineTimeStart := int(9e6)
		lineTimeEnd := 0
		lineIndex := 0
		for i, word := range textList {
			secondData := timestampList[i]
			if len(secondData) < 2 {
				continue
			}
			if secondData[0] < lineTimeStart {
				lineTimeStart = secondData[0]
			}
			if secondData[1] > lineTimeEnd {
				lineTimeEnd = secondData[1]
			}

			lineText = append(lineText, word)
			//if len(lineText) >= lineMaxLetter {
			//	srtList = append(srtList, fmt.Sprintf("%d --- %d\n%s",
			//		lineTimeStart,
			//		lineTimeEnd,
			//		strings.Join(lineText, ""),
			//	))
			//	// 重置
			//	lineTimeStart = 9e6
			//	lineTimeEnd = 0
			//	lineText = []string{}
			//	continue
			//}

			if i < textListLen-1 {
				//最后一个字
				if i == textListLen-2 {
					lineText = append(lineText, textList[textListLen-1])
				}
				if len(lineText) >= lineMaxLetter || timestampList[i+1][1]-secondData[1] >= lineBreakMillisecond {
					lineIndex++
					srtList = append(srtList, fmt.Sprintf("%d\n%s --> %s\n%s",
						lineIndex,
						lib.MillisecondsToVideoTime(int64(lineTimeStart)),
						lib.MillisecondsToVideoTime(int64(lineTimeEnd)),
						strings.Join(lineText, ""),
					))
					// 重置
					lineTimeStart = 9e6
					lineTimeEnd = 0
					lineText = []string{}
					continue
				}
			}
		}

	}

	return strings.Join(srtList, "\n\n")
}
