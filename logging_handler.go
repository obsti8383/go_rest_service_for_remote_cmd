package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingHandler struct {
	writer  io.Writer
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	t := time.Now().String()
	url := req.URL.EscapedPath()
	method := req.Method
	var logOut = fmt.Sprintf("%s %s %s\n", t, method, url)
	io.WriteString(h.writer, logOut)
	h.handler.ServeHTTP(res, req)
}

func LoggingHandler(out io.Writer, h http.Handler) http.Handler {
	return loggingHandler{out, h}
}
