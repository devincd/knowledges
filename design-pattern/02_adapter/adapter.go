// 适配器模式
// 适配器模式用于转换一种接口适配另一种接口
// 实际使用中Adaptee一般为接口，并且使用工厂函数生成实例
// 在Adapter中匿名组合Adaptee接口，所有Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非侵入式接口特征，其实Adapter也适配Adaptee接口
// 参考链接 http://c.biancheng.net/view/1361.html

package adapter

// Adaptee是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// adapteeImpl是被适配的目标类
type adapteeImpl struct{}

// SpecificRequest是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// Target是适配的目标接口
type Target interface {
	Request() string
}

// NewAdapter是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// adapter是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

// Request实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
