package designPattern

import (
	"context"
	"testing"
)

func TestFK(t *testing.T) {
	obA := NewBaseObserver("A")
	obA.Onchange(context.Background(), &Event{Topic: "哈哈哈", Value: 1})
}
