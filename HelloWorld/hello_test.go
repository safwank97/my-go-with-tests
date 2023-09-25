package main

import "testing"

/*func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World !!!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
*/

/*
func TestHelloGreetings(t *testing.T) {

	//Purposeful failing test with a name and not the string world
	got := HelloGreetings("Safwan !!!")
	want := "Hello, Safwan !!!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
*/

func TestHelloWithConstants(t *testing.T) {

	//Purposeful failing test with empty string and handling that,
	//while making changes to the .go file by passing a string under if string is empty then what to do
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWithConstatnts("Safwan !!!")
		want := "Hello, Safwan !!!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello World !!!' when an empty string is supplied ", func(t *testing.T) {
		got := HelloWithConstatnts("")
		want := "Hello, World !!!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
