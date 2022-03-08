// 包，表明代码所在的模块（包），程序入口必须这个包
package main

// 引入代码依赖
import (
	"fmt"
	"os"
)

// 功能实现
func main() {
	if len(os.Args) > 1 {
		fmt.Println("命令行参数不为空，", os.Args[1])
	}
	fmt.Println("好家伙，就是不行。")
	// 异常退出，控制台显示出了退出 code 255
	os.Exit(-1)
}
