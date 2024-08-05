package main

import (
	"fmt"
	component_shipinlv "github.com/tangwenru/go-component-shipinlv"
)

func main() {
	audioToText := component_shipinlv.AudioToText{}
	result, err := audioToText.Create(1, &component_shipinlv.AudioToTextCreateQuery{
		TaskId:        1,
		TaskType:      "audio-to-text",
		AudioFilePath: "https://shipinlv.oss-cn-hangzhou.aliyuncs.com/_tmp/9.mp3", // "https://isv-data.oss-cn-hangzhou.aliyuncs.com/ics/MaaS/ASR/test_audio/vad_example.wav",
	})

	fmt.Println("result:", result)
	fmt.Println("err:", err)
}
