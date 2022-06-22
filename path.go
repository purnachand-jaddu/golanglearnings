package main

import (
	"fmt"
	"runtime"

	"echoframework.com/basics/temp"
)

func main() {
	_, filename, _, _ := runtime.Caller(1)
	fmt.Println(filename)
	temp.Temp()
}
