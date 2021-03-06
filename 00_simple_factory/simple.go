package simplefactory

import "fmt"

// API is factory interface. Used as constructed obj
type API interface {
	Say(name string) string
}

// NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{} // todo: funny it returns address ?
	} else if t == 2 {
		return &helloAPI{} // todo: notice that address been used as interface type!!
	}
	return nil
}

// hiAPI is one of API implement
type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// HelloAPI is another API implement
type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
