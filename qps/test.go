package main

import (
	"fmt"
	"time"
)

// 工厂
type product interface {
	use() string
}
type productA struct{}

func (p *productA) use() string {
	return "生产A"
}

type productB struct{}

func (p *productB) use() string {
	return "生产B"
}
func Chooseproduct(chance string) product {
	switch chance {
	case "A":
		return &productA{}
	case "B":
		return &productB{}
	default:
		return nil
	}
}
func main() {
	for i := 0; i < 5; i++ {
		msg := "我是go0"
		go func(data string) { // 匿名函数 + 参数传递
			fmt.Println("消息:", data)
		}(msg)
		go go1()
		go go2("go1协程")
		time.Sleep(100 * time.Millisecond)
	}

}
func go1() {
	fmt.Printf("我是go1")
}
func go2(string1 string) {
	fmt.Printf("我是go2,%s", string1)
}
