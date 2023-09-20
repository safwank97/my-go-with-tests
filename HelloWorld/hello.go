package main

import "fmt"

/*
func Hello() string {
	return "Hello, world !!!"
}
*/

func HelloGreetings(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(HelloGreetings("Safwan !!!"))
	//fmt.Println(Hello())
}
