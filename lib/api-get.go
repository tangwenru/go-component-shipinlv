package lib

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
	"regexp"
	"strings"
)

func MainSystem(userId int64, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("ShiPinLv.APIUrl")
	return ApiGet(userId, apiUrl, apiPath, data, result)
}

func SubsystemChatgpt(userId int64, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("Subsystem.ChatgptAPIUrl")
	if len(apiUrl) < 5 {
		return nil, errors.New("请配置 app.conf 文件的 Subsystem.ChatgptAPIUrl")
	}
	return ApiGet(userId, apiUrl, apiPath, data, result)
}

func SubsystemChatgptApi(userId int64, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("Subsystem.ChatgptApiAPIUrl")
	if len(apiUrl) < 5 {
		return nil, errors.New("请配置 app.conf 文件的 Subsystem.ChatgptApiAPIUrl")
	}
	return ApiGet(userId, apiUrl, apiPath, data, result)
}

func SubsystemLongVideoSales(userId int64, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("Subsystem.LongVideoSalesAPIUrl")
	if len(apiUrl) < 5 {
		return nil, errors.New("请配置 app.conf 文件的 Subsystem.LongVideoSalesAPIUrl")
	}
	return ApiGet(userId, apiUrl, apiPath, data, result)
}

func SubsystemVideoPublish(userId int64, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("Subsystem.VideoPublishAPIUrl")
	if len(apiUrl) < 5 {
		return nil, errors.New("请配置 app.conf 文件的 Subsystem.VideoPublishAPIUrl")
	}
	return ApiGet(userId, apiUrl, apiPath, data, result)
}

func ApiGet(userId int64, apiUrl, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	//apiUrl := beego.AppConfig.String("ShiPinLv.APIUrl")
	url := apiUrl + strings.ToLower(apiPath)
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	dataByte, err := json.Marshal(data)

	postData := Encrypt(string(dataByte))
	req.Param("data", postData)
	req.Param("userId", IdEncrypt(userId))

	// 子系统名称
	appName, _ := beego.AppConfig.String("appName")
	reg := regexp.MustCompile("^api-")
	subSystem := reg.ReplaceAllString(appName, "")
	req.Param("subSystem", subSystem)

	//fmt.Println("ApiGet:", global.IdEncrypt(userId), url, postData)

	bytesResult, err := req.Bytes()
	if err != nil {
		fmt.Println("get api :", err)
		return []byte(""), err
	}

	//query := ThirdIpResult{}
	//errJson := json.Unmarshal(bytesResult, &query)
	//if errJson != nil {
	//	fmt.Println("err 21:", errJson)
	//	return ""
	//}

	errJson := json.Unmarshal(bytesResult, result)
	if errJson != nil {
		fmt.Println("err 1:", url)
		fmt.Println("err 21:", errJson, string(bytesResult))
		return []byte(""), errJson
	}

	return bytesResult, nil
}
