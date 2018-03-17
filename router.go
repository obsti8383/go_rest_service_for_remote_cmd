package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
// from https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)

	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type Router struct {
	// We could use http.Handler as a type here; using the specific type has
	// the advantage that static analysis tools can link directly from
	// h.UserHandler.ServeHTTP to the correct definition. The disadvantage is
	// that we have slightly stronger coupling. Do the tradeoff yourself.
	UserHandler         http.Handler
	RemoteAccessHandler http.Handler
}

func (h *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "user":
		h.UserHandler.ServeHTTP(res, req)
	case "remote":
		h.RemoteAccessHandler.ServeHTTP(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

////////// UserHandler
type UserHandler struct {
}

func (h *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Username = ", req.Context().Value("username"))

	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
		return
	}

	var response string

	switch req.Method {
	case "GET":
		response = h.handleGet(id)
	case "PUT":
		response = h.handlePut(id)
	default:
		http.Error(res, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(response); err != nil {
		panic(err)
	}
	//io.WriteString(res, response)
}

func (h *UserHandler) handleGet(id int) (response string) {
	return "hallo"
}

func (h *UserHandler) handlePut(id int) (response string) {
	return "hello"
}
