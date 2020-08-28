package leaf

import (
	"base/leaf/cluster"
	"base/leaf/conf"
	"base/leaf/console"
	"base/leaf/log"
	"base/leaf/module"
	"os"
	"os/signal"
)

func Run(mods ...module.Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Leaf %v starting up", version)

	// 注册模块 game login gate
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}

	// 模块管理器初始化，运行模块game login gate的Run函数
	module.Init()

	// cluster
	cluster.Init()

	// 初始化控制台模块
	console.Init()

	// 创建关闭通道变量，缓冲值为1，通道中无值时阻塞
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	sig := <-c

	log.Release("Leaf closing down (signal: %v)", sig)

	// 关闭通道接收道值，继续运行以下函数。倒序销毁各个模块
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
