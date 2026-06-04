package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAje(t *testing.T) {
	result := helloWorld("aje")

	if result != "hello aje" {
		// error
		t.Error("Result must be 'hello aje'")
	}
	fmt.Println("TestHelloWorldAje Done")
}

func TestHelloWorldAndhika(t *testing.T) {
	result := helloWorld("andhika")

	if result != "hello andhika" {
		// error
		t.Fatal("Result must be 'hello andhika'")
	}
	fmt.Println("TestHelloWorldAndhika Done")
}
