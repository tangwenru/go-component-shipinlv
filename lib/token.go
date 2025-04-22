package lib

import (
	"fmt"
	"time"
)

func MakeToken(aesKey string) string {
	randNum1 := RandWord(5, "")
	randNum2 := RandWord(5, "")

	text := fmt.Sprintf("%s_%d_%s",
		randNum1,
		time.Now().Unix()+3600,
		randNum2,
	)
	return EncryptText(text, aesKey)
}

func CheckToken(aesKey, token string) bool {
	text := DecryptText(token, aesKey)
	textData := Split(text, "_")
	if len(textData) <= 1 || textData[1] == "" {
		return false
	}

	expired := String2Int64(textData[1])

	return expired > time.Now().Unix()
}
