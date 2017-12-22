package test

import (
	"io"
	"net/http"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>Hello World!</body></html>")
}
