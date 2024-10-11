package concurrency

import "testing"

func TestConcurrencyAssignment1(t *testing.T) {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}
