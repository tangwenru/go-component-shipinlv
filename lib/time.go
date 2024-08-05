package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// MillisecondsToVideoTime 将毫秒数转换为视频时长格式（小时:分钟:秒.毫秒）
func MillisecondsToVideoTime(milliseconds int64) string {
	hours := milliseconds / (1000 * 60 * 60)
	remainingMilliseconds := milliseconds % (1000 * 60 * 60)
	minutes := remainingMilliseconds / (1000 * 60)
	remainingMilliseconds = remainingMilliseconds % (1000 * 60)
	seconds := remainingMilliseconds / 1000
	millisecondsPart := remainingMilliseconds % 1000

	// 确保毫秒部分为三位数
	millisecondsStr := strconv.FormatInt(millisecondsPart, 10)
	if len(millisecondsStr) < 3 {
		millisecondsStr = strings.Repeat("0", 3-len(millisecondsStr)) + millisecondsStr
	}

	return fmt.Sprintf("%d:%02d:%02d.%s", hours, minutes, seconds, millisecondsStr)
}
