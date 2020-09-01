package internal

import (
	"base/leaf/gate"
	//"base/leaf/log"
	"fmt"
	"reflect"
	"server/msg"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.HelloT{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.HelloT)

	// 消息的发送者
	a := args[1].(gate.Agent)

	// 给发送者回应一个 Hello 消息
	m.Name = "liu"
	a.WriteMsg(m)

	// 输出收到的消息的内容
	fmt.Println(m)
}