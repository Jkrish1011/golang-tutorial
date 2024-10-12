package concurrency

import (
	"fmt"
	"testing"
)

func TestConcurrencyAssignment1(t *testing.T) {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}

func TestSampleConcurrency(t *testing.T) {
	ch := make(chan bool)
	go sampleFunctionForConcurrency("Hello", ch)
	println(<-ch)
}

func TestChannel2(t *testing.T) {
	ch := make(chan bool, 2)
	ch <- true
	ch <- false
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
