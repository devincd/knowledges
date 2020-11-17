// 简单工厂模式
// go语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。NewXXX函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂
// 在这个simplefactory包中只有API接口和NewAPI函数为包外可见，封装了实现细节
// 参考链接 http://c.biancheng.net/view/8385.html

package simplefactory

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI return API instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI is one of API implement
type hiAPI struct{}

// Check that hiAPI implements the API interface
var _ API = &hiAPI{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI is another API implement
type helloAPI struct{}

// say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
