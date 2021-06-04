package internal

import (
	"base/leaf/gate"
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
	"server/data/role"
	"server/msg"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Sync{}, handleSync)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleSync(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Sync)
	// 消息的发送者
	a := args[1].(gate.Agent)

	data := a.UserData().(*role.Role)
	data.X = *m.X
	data.Y = *m.Y
	Broadcast()

	// 输出收到的消息的内容
	fmt.Println(*data)
}

// 同步给所有玩家，所有玩家的位置
func Broadcast() {
	roles := msg.Broadcast{}
	for a := range agents {
		data := a.UserData().(*role.Role)
		roles.Roles = append(roles.Roles, &msg.Sync{
			Id: proto.Int32(data.Id),
			X: proto.Int32(data.X),
			Y: proto.Int32(data.Y),
		})
	}

	for a := range agents {
		a.WriteMsg(&roles)
	}
}