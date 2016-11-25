package main

import "fmt"
import "runtime"

func main() {
	fmt.Println("Hello,", runtime.GOARCH, "world!")
}
