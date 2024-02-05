package subsystemChatgptApi

import (
	"fmt"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
	typeChatgptApi "github.com/tangwenru/go-component-shipinlv/subsystem/chatgpt-api/type"
)

type Vip struct {
}

func init() {

}

// @VersionCategory ： 3.5 | 4.0 ...
func (this *Vip) PayedUpdate(userId int64, versionCategory string, dollarAmount float64) (error, *typeChatgptApi.VipPayedUpdateData) {
	query := typeChatgptApi.VipPayedUpdateQuery{
		VersionCategory: versionCategory,
		DollarAmount:    dollarAmount,
	}
	vipPayedUpdate := typeChatgptApi.VipPayedUpdateData{}

	bytesResult, err := component_shipinlv_lib.SubsystemChatgptApi(userId, "vip/PayedUpdate", &query, &vipPayedUpdate)

	if err != nil {
		fmt.Println("SubsystemChatgptApi) Detail:", string(bytesResult))
	}

	return err, &vipPayedUpdate
}
