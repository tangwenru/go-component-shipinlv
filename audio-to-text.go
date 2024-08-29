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

type AudioToText struct {
}

type AudioToTextDetail struct {
	Id           int64  `json:"id"`
	WorkServerId int64  `json:"workServerId"`
	SrtText      string `json:"srtText"`
}

type AudioToTextDetailResult struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    AudioToTextDetail `json:"data"`
}

type AudioToTextCreateQuery struct {
	TaskId        int64  `json:"taskId"`
	TaskType      string `json:"taskType"`
	LineMaxLetter int    `json:"lineMaxLetter"`
	AudioFilePath string `json:"audioFilePath"`
}

type AudioToTextApiResult struct {
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

func (this *AudioToText) Create(userId int64, query *AudioToTextCreateQuery) (*AudioToTextDetail, error) {
	// 先找 一个 服务器
	workServer := WorkServer{}
	workServerDetail, errWorkServerDetail := workServer.RandDetail(userId, "audio-to-text")
	if errWorkServerDetail != nil {
		return nil, errWorkServerDetail
	}

	port := fmt.Sprintf("%d", workServerDetail.Port)
	if len(port) > 1 {
		port = ":" + port
	}
	apiUrl := fmt.Sprintf("%s://%s%s/api-audioToText/audioToText",
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
		fmt.Println("get api:", err)
		return nil, errors.New("识别字幕失败 0：" + err.Error())
	}

	resultData := AudioToTextApiResult{}
	json.Unmarshal(byteResult, &resultData)

	if !resultData.Success {
		return nil, errors.New("识别字幕失败 1：" + resultData.Message)
	}

	fmt.Println("api:", resultData)
	outDetail := AudioToTextDetail{
		WorkServerId: workServerDetail.Id,
		SrtText:      this.ApiResult2Srt(query.LineMaxLetter, &resultData),
	}

	return &outDetail, nil
}

func (this *AudioToText) ApiResult2Srt(lineMaxLetter int, data *AudioToTextApiResult) string {
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
		// 文本 太长
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
