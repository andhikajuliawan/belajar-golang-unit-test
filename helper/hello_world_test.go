package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := helloWorld("aje")

	assert.Equal(t, "hello aje", result, "Result must be 'hello aje'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := helloWorld("aje")

	require.Equal(t, "hello aje", result, "Result must be 'hello aje'")
	fmt.Println("TestHelloWorld with Require Done")
}

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
