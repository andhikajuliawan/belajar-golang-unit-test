package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := helloWorld("aje")

	if result != "hello aje" {
		// error
		panic("Result is not hello aje")
	}

}

func TestHelloWorldAndhika(t *testing.T) {
	result := helloWorld("andhika")

	if result != "hello andhika" {
		// error
		panic("Result is not hello andhika")
	}
}
