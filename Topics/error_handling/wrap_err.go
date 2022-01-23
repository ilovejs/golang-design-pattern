package error_handling

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// 在 Go 语言的开发者中，更为普遍的做法是将错误包装在另一个错误中，同时保留原始内容

type authorizationError struct {
	operation string
	err       error // original error
}

func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

// 当然，更好的方式是通过一种标准的访问方法，
// 这样，我们最好使用一个接口，比如 causer接口中实现 Cause() 方法来暴露原始错误，以供进一步检查

type causer interface {
	Cause() error
}

func (e *authorizationError) Cause() error {
	return e.err
}

type MyError struct {
	a int
}

func (e *MyError) Error() string {
	return "MyError"
}

func UseThirdParty() (err error) {

	// 错误包装
	if err != nil {
		// 3rd party lib
		return errors.Wrap(err, "read failed")
	}

	// Cause接口
	switch err := errors.Cause(err).(type) {
	case *MyError:
		// handle specifically
		log.Fatalf("MyError %s", err)
	default:
		log.Fatalf("unknown error %s", err)
	}

	return
}
