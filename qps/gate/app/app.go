package app

import (
	"context"
	"fmt"
	"gate/router"
	"msqp/config"
	"msqp/logs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context) error {
	// 初始化日志库
	logs.InitLog(config.Conf.AppName)
	go func() {
		r := router.RegisterRouter()
		// http接口
		if err := r.Run(fmt.Sprintf(":%d", config.Conf.HttpPort)); err != nil {
			logs.Fatal("gate gin run err:%v", err)
		}
	}()

	stop := func() {
		time.Sleep(3 * time.Second)
		fmt.Println("stop app finish")
	}

	// 期望优雅启停 遇到中断，退出，终止，挂断
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
	for {
		select {
		case <-ctx.Done():
			stop()
			return nil
		case s := <-c:
			switch s {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
				stop()
				logs.Info("user app quit")
			case syscall.SIGHUP:
				stop()
				logs.Info("hang up!! user app quit")
				return nil
			default:
				return nil
			}
		}
	}
}
