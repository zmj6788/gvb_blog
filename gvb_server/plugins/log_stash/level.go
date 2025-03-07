package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel Level = 1 // 调试级别
	InfoLevel  Level = 2 // 普通级别
	WarnLevel  Level = 3 // 警告级别
	ErrorLevel Level = 4 // 错误级别
)

// 角色json序列化
func (r Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// 角色匹配
func (r Level) String() string {
	switch r {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "其他"
	}
}
