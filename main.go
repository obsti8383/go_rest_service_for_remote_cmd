package main

import "net/http"
import "os"

func main() {
	a := &Router{
		UserHandler:         AuthenticationHandler(LoggingHandler(os.Stdout, new(UserHandler))),
		RemoteAccessHandler: AuthenticationHandler(LoggingHandler(os.Stdout, new(RemoteAccessHandler))),
	}
	http.ListenAndServe(":8000", a)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
