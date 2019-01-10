package minstack

import "testing"

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	if res := obj.GetMin(); res != -3 {
		t.Errorf("expected %d, got %d", -3, res)
	}

	obj.Pop()
	if res := obj.Top(); res != 0 {
		t.Errorf("expected %d, got %d", 0, res)
	}
	if res := obj.GetMin(); res != -2 {
		t.Errorf("expected %d, got %d", -2, res)
	}
}
