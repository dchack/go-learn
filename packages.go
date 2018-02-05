package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	fmt.Print("not hello world", rand.Intn(10))
	//在导入包之后，你只能访问包所导出的名字，任何未导出的名字是不能被包外的代码访问的。
	fmt.Println(math.Pi)
}
