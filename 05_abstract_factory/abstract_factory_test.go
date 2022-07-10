package abstractfactory

func ExampleRdbFactory() {
	var factory DAOFactory

	factory = &RDBDAOFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()

	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXmlFactory() {
	var factory DAOFactory

	factory = &XMLDAOFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()

	// Output:
	// xml main save
	// xml detail save
}

func Example() {
	dao := new(RDBDAOFactory)
	dao.CreateOrderMainDAO().SaveOrderMain()
	dao.CreateOrderDetailDAO().SaveOrderDetail()
}

func Example2() {
	dao := new(XMLDAOFactory)
	dao.CreateOrderMainDAO().SaveOrderMain()
	dao.CreateOrderDetailDAO().SaveOrderDetail()
}
