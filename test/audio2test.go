package main

import (
	"fmt"
	component_shipinlv "github.com/tangwenru/go-component-shipinlv"
)

func main() {
	audio2Text := component_shipinlv.Audio2Text{}
	result, err := audio2Text.Create(1, &component_shipinlv.Audio2TextCreateQuery{
		TaskId:        1,
		TaskType:      "audio2text",
		AudioFilePath: "https://shipinlv.oss-cn-hangzhou.aliyuncs.com/_tmp/9.mp3", // "https://isv-data.oss-cn-hangzhou.aliyuncs.com/ics/MaaS/ASR/test_audio/vad_example.wav",
	})

	fmt.Println("result:", result)
	fmt.Println("err:", err)
}
