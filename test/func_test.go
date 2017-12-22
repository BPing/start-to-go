package test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

//func TestFoo（）//测试Foo函数或类型
//func TestBar_Qux（）//测试Bar类型的Qux方法
//func Test（）//将软件包作为一个整体文档

func TestFibonacci(t *testing.T) {
	r := Fibonacci(10)
	if r != 55 {
		t.Errorf("Fibonacci(10) failed. Got %d, expected 55.", r)
	}
}

func TestHttpHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	HttpHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		t.Errorf("StatusCode failed. Got %d, expected 200.", resp.StatusCode)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
