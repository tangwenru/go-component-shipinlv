package lib

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/shopspring/decimal"
)

type DataResultModel struct {
	Success    bool        `json:"success"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	ServerTime int64       `json:"serverTime"`
	Data       interface{} `json:"data"`
}

func GetDataResultModel() DataResultModel {
	result := DataResultModel{
		Code:       "0",
		Success:    false,
		ServerTime: time.Now().Unix(),
	}
	return result
}

func Json2String(data interface{}) string {
	resultStr, errJson := json.Marshal(&data)
	if errJson != nil {

	}
	return string(resultStr)
}

func DataResultSuccess(message string, data interface{}) (r DataResultModel) {
	result := DataResultModel{}

	result.Success = true
	result.Code = "0"
	result.Message = message
	result.Data = data
	result.ServerTime = time.Now().Unix()
	return result
}

func DataResultError(code string, message string, data interface{}) (r DataResultModel) {
	result := DataResultModel{}

	result.Success = code == "0"
	result.Code = code
	result.Message = message
	result.Data = data
	result.ServerTime = time.Now().Unix()
	return result
}

// 截取字符串
func Slice(word string, startPos int, endPos int) string {
	wordRune := []rune(word)

	wordLen := len(wordRune)

	max := Min(int64(wordLen), int64(endPos))
	if max < int64(startPos) {
		max = int64(wordLen)
	}

	fmt.Println("wordLen:", wordLen, max, endPos)
	fmt.Println("wordLen startPos:", startPos)

	if startPos >= wordLen {
		return ""
	}
	return string(wordRune[startPos:max])
}

// 四舍五入
func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

// 四舍五入
func RoundToFloat64(x float64) float64 {
	return float64(int64(math.Floor(x + 0.5)))
}

func String36To10(string36 string) int64 {
	i, err := strconv.ParseUint(string36, 36, 64)
	if err != nil {
		return 0
	}
	return int64(i)
}

func String2Int64(str string) int64 {
	if str == "" {
		return 0
	}

	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		id = 0
	}
	return id
}

func String2Int8(str string) int8 {
	if str == "" {
		return 0
	}

	id, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		id = 0
	}
	return int8(id)
}

func Float64ToString(num float64) string {
	return strconv.FormatFloat(num, 'g', -1, 64)
}

func Int64ToString(int int64) string {
	out := strconv.FormatInt(int, 10)
	return out
}
func IntToString(int int) string {
	out := strconv.Itoa(int)
	return out
}

func String2Float64(str string) float64 {
	id, err := strconv.ParseFloat(str, 64)
	if err != nil {
		id = 0
	}
	return id
}

func String13ToTime10(str string) int64 {
	return String2Int64(string([]rune(str)[:10]))
}

func Strings2Int64(str string) []int64 {
	tradeIdsStringArray := Split(str, ",")
	var ids []int64
	for _, idString := range tradeIdsStringArray {
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			id = 0
		}
		ids = append(ids, id)
	}
	return ids
}

func String2Array(str string, split string) []string {
	if split == "" {
		split = ","
	}

	if str == "" {
		return []string{}
	}

	return Split(str, split)
}

// 100-200 转换 数字区间
func String2NumberArea(str string, maxNumber int64) (int64, int64) {
	arrayData := String2Array(str, "-")
	// 全部；
	if len(arrayData) != 2 {
		return 0, maxNumber
	}

	min := String2Int64(arrayData[0])
	max := String2Int64(arrayData[1])
	if max == 0 {
		max = maxNumber
	}

	return min, max
}

func Cent2Yen(money int64) float64 {
	x := decimal.NewFromFloat(float64(money))
	var z = x.Div(decimal.NewFromInt32(100))
	moneyYen, _ := strconv.ParseFloat(z.StringFixed(2), 64)
	return moneyYen
}

func Yen2Cent(money float64) int64 {
	return int64(money * 100)
}

// float64 四目运算；
// 加
func Add(a float64, b float64) float64 {
	x := decimal.NewFromFloat(a)
	y := decimal.NewFromFloat(b)

	sum := x.Add(y)

	result, err := strconv.ParseFloat(sum.StringFixed(16), 64)
	if err != nil {
		result = 0
	}
	return result
}

// 减
func Sub(a float64, b float64) float64 {
	x := decimal.NewFromFloat(a)
	y := decimal.NewFromFloat(b)

	sum := x.Sub(y)

	result, err := strconv.ParseFloat(sum.StringFixed(16), 64)
	if err != nil {
		result = 0
	}
	return result
}

// 乘
func Mul(a float64, b float64) float64 {
	x := decimal.NewFromFloat(a)
	y := decimal.NewFromFloat(b)

	sum := x.Mul(y)

	result, err := strconv.ParseFloat(sum.StringFixed(16), 64)
	if err != nil {
		result = 0
	}
	return result
}

// 除
func Div(a, b float64) float64 {
	aa := decimal.NewFromFloat(a)
	var bb = aa.Div(decimal.NewFromFloat(b))
	result, err := strconv.ParseFloat(bb.StringFixed(16), 64)
	if err != nil {
		result = 0
	}
	return result
}

// id 的 大于 0；
func IdEncrypt(id int64) string {
	if id <= 0 {
		return ""
	}
	aesKeyString, _ := beego.AppConfig.String("AesKey")
	var aesKey = []byte(aesKeyString)
	result, err := AesEncrypt([]byte(Int64ToString(id)), aesKey)
	if err != nil {
		return ""
	}
	return result
}

func IdDecrypt(idKey string) int64 {
	aesKeyString, _ := beego.AppConfig.String("AesKey")
	var aesKey = []byte(aesKeyString)
	result, err := AesDecrypt(idKey, aesKey)
	if err != nil {
		return 0
	}
	return String2Int64(string(result))
}

func Encrypt(text string) string {
	return EncryptText(text, "")
}

func Decrypt(idKey string) string {
	return DecryptText(idKey, "")
}

func EncryptText(text string, aesKey string) string {
	if text == "" {
		return ""
	}
	aesKeyString, _ := beego.AppConfig.String("AesKey")
	if aesKey == "" {
		aesKey = aesKeyString
	}
	var aesKeyByte = []byte(aesKey)
	result, err := AesEncrypt([]byte(text), aesKeyByte)
	if err != nil {
		return ""
	}
	return result
}
func DecryptText(idKey string, aesKey string) string {
	aesKeyString, _ := beego.AppConfig.String("AesKey")
	if aesKey == "" {
		aesKey = aesKeyString
	}
	var aesKeyByte = []byte(aesKeyString)
	result, err := AesDecrypt(idKey, aesKeyByte)
	if err != nil {
		return ""
	}
	return string(result)
}

func RandWord(len int, words string) string {
	maxLen := len
	var container string
	var strDict = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	if words != "" {
		strDict = words
	}
	b := bytes.NewBufferString(strDict)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < maxLen; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(strDict[randomInt.Int64()])
	}
	return container
}

// 整数
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func Split(text, sep string) []string {
	if text == "" {
		return []string{}
	}
	return strings.Split(text, sep)
}
