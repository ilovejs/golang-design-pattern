package facade

import "fmt"

// API is shared methods among sub-system
type API interface {
	Test() string
}

// NewAPI factory method
func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// ---------------------------------

// apiImpl is impl of API interface
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

// Test belongs to the root
func (a *apiImpl) Test() string {
	// calls test in sub-system
	aRet := a.a.TestAXX()
	bRet := a.b.TestB()

	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// ---------------------------------

// NewAModuleAPI simple factory
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type AModuleAPI interface {
	TestAXX() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestAXX() string {
	return "A module running"
}

// ---------------------------------

// NewBModuleAPI simple factory return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
