package main

import (
	"connector/app"
	"context"
	"fmt"
	"framework/game"
	"github.com/spf13/cobra"
	"log"
	"msqp/config"
	"msqp/metrics"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "connector",
	Short: "connector 管理连接，session以及路由请求",
	Long:  `connector 管理连接，session以及路由请求`,
	Run: func(cmd *cobra.Command, args []string) {
	},
	PostRun: func(cmd *cobra.Command, args []string) {
	},
}

// var configFile = flag.String("config", "application.yml", "config file")
var (
	configFile    string
	gameConfigDir string
	serverId      string
)

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", "application.yml", "app config yml file")
	rootCmd.Flags().StringVar(&gameConfigDir, "gameDir", "../config", "game config dir")
	rootCmd.Flags().StringVar(&serverId, "serverId", "", "app server id， required")
	_ = rootCmd.MarkFlagRequired("serverId")
}
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	game.InitConfig(gameConfigDir)
	config.InitConfig(configFile)
	fmt.Println(config.Conf)
	//2.启动监控
	go func() {
		err := metrics.Serve(fmt.Sprintf("0.0.0.0:%d", config.Conf.MetricPort))
		if err != nil {
			panic(err)
		}
	}()
	err := app.Run(context.Background(), serverId)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

}
