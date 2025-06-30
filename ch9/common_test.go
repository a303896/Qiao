package ch9

import "testing"

// 单线程测试
func TestGetContent(t *testing.T) {
	//m := New(httpGetBody)
	m := NewMemo2(httpGetBody)
	getContent(m)
}

// 多线程测试
func TestGetContentMulti(t *testing.T) {
	//m := New(httpGetBody)
	m := NewMemo2(httpGetBody)
	getContentMulti(m)
}
