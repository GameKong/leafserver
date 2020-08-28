package module

import (
	"base/leaf/conf"
	"base/leaf/log"
	"runtime"
	"sync"
)

type Module interface {
	OnInit()
	OnDestroy()
	Run(closeSig chan bool)
}

type module struct {
	mi       Module
	closeSig chan bool
	wg       sync.WaitGroup
}

var mods []*module

func Register(mi Module) {
	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool, 1)

	// 创建模块加入到模块列表中
	mods = append(mods, m)
}

// 模块管理器
func Init() {
	// 执行模块的OnInit初始化函数
	for i := 0; i < len(mods); i++ {
		mods[i].mi.OnInit()
	}

	for i := 0; i < len(mods); i++ {
		m := mods[i]
		m.wg.Add(1)
		go run(m)
	}
}

// 主函数调用
func Destroy() {
	for i := len(mods) - 1; i >= 0; i-- {
		m := mods[i]
		m.closeSig <- true // 给予关闭信号
		m.wg.Wait() // 阻塞 等到等待组中的所有任务结束后再进行进行
		destroy(m)
	}
}

func run(m *module) {
	m.mi.Run(m.closeSig) // skeleton的Run函数，执行select，执行相应通道对应的事件函数
	m.wg.Done()
}

func destroy(m *module) {
	// 函数结束前调用
	defer func() {
		if r := recover(); r != nil {
			if conf.LenStackBuf > 0 {
				buf := make([]byte, conf.LenStackBuf)
				l := runtime.Stack(buf, false)
				log.Error("%v: %s", r, buf[:l])
			} else {
				log.Error("%v", r)
			}
		}
	}()

	// 子模块的销毁
	m.mi.OnDestroy()
}
