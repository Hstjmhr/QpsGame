package app

import (
	"connector/route"
	"context"
	"core/repo"
	"fmt"
	"framework/connector"
	"msqp/config"
	"msqp/logs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context, serverId string) error {
	// 初始化日志库
	logs.InitLog(config.Conf.AppName)
	exit := func() {}
	go func() {
		c := connector.Default()
		exit = c.Close
		manager := repo.New()
		c.RegisterHandler(route.Register(manager))
		c.Run(serverId)
	}()

	stop := func() {
		exit()
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
				logs.Info("connector app quit")
			case syscall.SIGHUP:
				stop()
				logs.Info("hang up!! connector app quit")
				return nil
			default:
				return nil
			}
		}
	}
}
