package arrays

import (
	"fmt"
	"testing"
)

func TestArrayExample(t *testing.T) {
	arrayExample()
}

func TestAssignment1(t *testing.T) {
	messages, costs := getMessageWithRetries("hello", "world", "!")
	fmt.Printf("The messages are %+v and costs are %+v", messages, costs)
}
