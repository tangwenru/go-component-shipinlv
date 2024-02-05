package lib

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

func MainSystem(userId int64, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("KTVAI.APIUrl")
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

func ApiGet(userId int64, apiUrl, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	//apiUrl := beego.AppConfig.String("KTVAI.APIUrl")
	url := apiUrl + strings.ToLower(apiPath)
	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	dataByte, err := json.Marshal(data)

	postData := Encrypt(string(dataByte))
	req.Param("data", postData)
	req.Param("userId", IdEncrypt(userId))

	//fmt.Println("ApiGet:", global.IdEncrypt(userId), url, postData)

	bytesResult, err := req.Bytes()
	if err != nil {
		fmt.Println("get ip :", err)
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
		fmt.Println("err 21:", errJson)
		return []byte(""), errJson
	}

	return bytesResult, nil
}
