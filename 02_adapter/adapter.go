package adapter

// Target 是适配的目标接口, 即适配器的实现目标是一个接口 [1]
type Target interface {
	// Request 会调用底层实现，而实现的函数也来自接口
	Request() string
}

// Adaptee 是被适配的目标接口 [3]
type Adaptee interface {
	SpecificRequest() string
}

//-- Adapter --//

// Adapter 适配器,转换Adaptee适合Target接口 [4]
type adapter struct {
	Adaptee
}

// NewAdapter 工厂函数
// todo: 返回一个接口，所以参考接口的实现
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		// todo: 嵌入
		Adaptee: adaptee,
	}
}

// Request 实现Target接口 [2]
func (a *adapter) Request() string {
	// todo: 调用接口函数
	// 调用嵌入的adaptee的函数
	return a.SpecificRequest()
}

//-- Adaptee --//

// NewAdaptee 工厂函数
// todo: 返回的是接口，不是实现
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct{}

// SpecificRequest 满足了接口需求
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee specific method impl..."
}
