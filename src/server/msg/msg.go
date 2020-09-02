package msg

import (
	"base/leaf/network/protobuf"
)

// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	// 这里我们注册了消息 Hello
	Processor.Register(&Hello{})
	Processor.Register(&Hello2{})
	Processor.Register(&Hello3{})
	Processor.Register(&Hello4{})
}
