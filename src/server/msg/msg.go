package msg

import (
	"base/leaf/network/json"
)

// 使用 Protobuf 消息处理器
var Processor = json.NewProcessor()


func init() {
	// 这里我们注册了消息 Hello
	Processor.Register(&HelloT{})
}
