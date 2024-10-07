package subsystemVideoPublish

import (
	"fmt"
	component_shipinlv_lib "github.com/tangwenru/go-component-shipinlv/lib"
	typeVideoTopic "github.com/tangwenru/go-component-shipinlv/subsystem/video-publish/type"
)

type VideoTopic struct {
}

func init() {

}

func (this *VideoTopic) TopicList(
	userId int64,
	query *typeVideoTopic.VideoTopicTopicListQuery,
) (error, *[]typeVideoTopic.VideoTopicTopicList) {
	result := typeVideoTopic.VideoTopicListResult{}

	bytesResult, err := component_shipinlv_lib.SubsystemVideoPublish(
		userId,
		"system/videoTopic/list",
		query,
		&result,
	)

	if err != nil {
		fmt.Println("VideoTopic list :", string(bytesResult))
		return err, nil
	}

	return nil, &result.Data
}
