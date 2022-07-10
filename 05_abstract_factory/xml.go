package abstractfactory

import "fmt"

// XMLMainDAO XML存储
type XMLMainDAO struct{}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// XMLDetailDAO XML存储
type XMLDetailDAO struct{}

func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

// XMLDAOFactory 是RDB 抽象工厂实现 [5]
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
