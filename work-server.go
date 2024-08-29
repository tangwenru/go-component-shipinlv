package component_shipinlv

import (
	"errors"
	"fmt"
	//component_shipinlv_lib "component_shipinlv/lib"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
)

type WorkServer struct {
}

type WorkServerDetail struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	Domain         string `json:"domain"`
	DomainProtocol string `json:"domainProtocol"`
	Port           int64  `json:"port"`
	WorkType       string `json:"workType"`
}

type WorkServerDetailQuery struct {
	WorkType string `json:"workType"`
}

type WorkServerDetailResult struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    WorkServerDetail `json:"data"`
}

func init() {

}

func (this *WorkServer) RandDetail(userId int64, workType string) (*WorkServerDetail, error) {
	query := WorkServerDetailQuery{
		WorkType: workType,
	}
	vipListResult := WorkServerDetailResult{}

	bytesResult, err := component_shipinlv_lib.MainSystem(userId, "workServer/randDetail", &query, &vipListResult)

	if err != nil {
		fmt.Println("Work Server RandDetail err:", string(bytesResult), err)
	}

	if !vipListResult.Success {
		return nil, errors.New(vipListResult.Message)
	}

	return &vipListResult.Data, err
}
