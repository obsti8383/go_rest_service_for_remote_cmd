package main

import "net/http"
import "context"

const ContextUsernameKey = "username"

type authenticationHandler struct {
	handler http.Handler
}

// returns a new AuthenticationHandler that uses the http.Handler in argument h to serve
// the request if authenticated successfully
func AuthenticationHandler(h http.Handler) http.Handler {
	return authenticationHandler{h}
}

// verifies if user is authenticated via BasicAuth
// returns error 401, if not
// calls handler that was provided to AuthenticationHandler(), if yes
// put username (with key "username") into context for next handler
func (h authenticationHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	user, pass, _ := req.BasicAuth()
	if !check(user, pass) {
		res.Header().Set("WWW-Authenticate", "Basic realm=\"TestRealm\"")
		http.Error(res, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	// put username into context for next handler
	ctx := context.WithValue(req.Context(), ContextUsernameKey, user)
	// call next handler
	h.handler.ServeHTTP(res, req.WithContext(ctx))
}

// just a dummy right now checking for test/test
func check(user, pass string) bool {
	if user == "test" && pass == "test" {
		return true
	} else {
		return false
	}
}
