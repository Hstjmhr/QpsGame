package app

import (
	"context"
	"core/repo"
	"fmt"
	"google.golang.org/grpc"
	"msqp/config"
	"msqp/discovery"
	"msqp/logs"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user/internal/service"
	"user/pb"
)

func Run(ctx context.Context) error {
	// 初始化日志库
	logs.InitLog(config.Conf.AppName)
	// etcd注册中心 grpc服务注册到etcd中 客户端访问的时候 通过etcd获取grpc地址
	register := discovery.NewRegister()
	// 启动grpc服务
	server := grpc.NewServer()
	// 初始化数据库
	manager := repo.New()
	go func() {
		lis, err := net.Listen("tcp", config.Conf.Grpc.Addr)
		if err != nil {
			logs.Fatal("user grpc server listen err:%v", err)
		}
		// 在etcd中注册服务
		err = register.Register(config.Conf.Etcd)
		if err != nil {
			logs.Fatal("user grpc server register etcd err:%v", err)
		}

		pb.RegisterUserServiceServer(server, service.NewAccountService(manager))

		// 阻塞操作
		err = server.Serve(lis)
		if err != nil {
			logs.Fatal("user grpc server run fail err:%v", err)
		}
	}()

	stop := func() {
		server.Stop()
		register.Close()
		manager.Close()
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
