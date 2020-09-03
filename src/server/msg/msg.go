package msg

import (
	"base/leaf/network/protobuf"
)

// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	// 这里我们注册了消息 Hello
	Processor.Register(&Sync{}) // id:1
	Processor.Register(&Broadcast{}) // id:2
}
