package main

import "fmt"

func SampleFunctionA(s string) string {
	return s
}

func main() {
	var str = SampleFunctionA("twinki")
	fmt.Println(str)
}
