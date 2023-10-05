package main

import "fmt"

/*
func Hello() string {
	return "Hello, World !!!"	//Replace World with your name to understand difference better.
}
*/

/*
func HelloGreetings(name string) string {
	return "Hello, " + name
}
*/

const englishHelloPrefix = "Hello, "

func HelloWithConstatnts(name string) string {

	if name == "" {
		name = "World !!!"
	}

	return englishHelloPrefix + name
}

func main() {
	//fmt.Println(Hello())
	//fmt.Println(HelloGreetings("World !!!"))
	fmt.Println(HelloWithConstatnts("World !!!")) //Replace World with your name to understand difference better.
}
