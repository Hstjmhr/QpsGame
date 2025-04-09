package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/big_model_ser"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	msgChan, err := big_model_ser.Send("qwen", "讲一个笑话")
	if err != nil {
		fmt.Println(err)
		return
	}
	for msg := range msgChan {
		fmt.Println(msg)
	}

}
