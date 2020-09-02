package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 设置收到消息的处理模块， 处理模块为game.ChanRPC
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}