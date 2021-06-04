package internal

import (
	"base/leaf/module"
	"server/base"
)

// 导入模块时就创建了以下对象
var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}

func (m *Module) BroadCast(msg string) {
	for a := range agents {
		a.WriteMsg(msg)
	}
}
