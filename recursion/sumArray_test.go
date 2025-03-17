package recursion

import (
	"Qiao"
	"testing"
)

type testObj struct {
	params []int
	result int
}

type testObj2 struct {
	params []int
	result []int
}

func TestSumArray(t *testing.T) {
	tests := []testObj{
		{[]int{2, 4, 6}, 12},
	}
	for _, test := range tests {
		if r := SumArray(test.params); r != test.result {
			t.Errorf("param:%v expected:%v got:%v", test.params, test.result, r)
		}
	}
}

func TestArrayLength(t *testing.T) {
	tests := []testObj{
		{[]int{2, 4, 6}, 3},
	}
	for _, test := range tests {
		if r := ArrayLength(test.params); r != test.result {
			t.Errorf("param:%v expected:%v got:%v", test.params, test.result, r)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []testObj2{
		{[]int{10, 5, 2, 3}, []int{2, 3, 5, 10}},
	}
	for _, test := range tests {
		r := QuickSort(test.params)
		t.Logf("param:%v expected:%v got:%v\n", test.params, test.result, r)
		if !Qiao.Equal(r, test.result) {
			t.Errorf("param:%v expected:%v got:%v", test.params, test.result, r)
		}
	}
}
