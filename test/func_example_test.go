package test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
)

// Example:https://blog.golang.org/examples
// godoc   -http=:7171

func ExampleHttpHandler() {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	HttpHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}
