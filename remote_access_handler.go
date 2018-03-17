package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

type RemoteAccessHandler struct {
}

func (h *RemoteAccessHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Username = ", req.Context().Value("username"))

	var command, out string
	var err error

	command, req.URL.Path = ShiftPath(req.URL.Path)

	switch command {
	case "ls":
		out, err = executeCommand("ls", "-l")
	case "lsetc":
		out, err = executeCommand("ls", "-l", "/etc")
	case "ping":
		out, err = executeCommand("ping", "-c 5", "www.heise.de")
	case "ip":
		out, err = executeCommand("ip", "addr")
	default:
		res.WriteHeader(http.StatusBadRequest)
		//io.WriteString(res, err.Error())
		return
	}

	if err != nil {
		//res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(res, err.Error())
	} else {
		//res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusOK)
		io.WriteString(res, out)
	}

}

// helper method for executing commands on OS
func executeCommand(cmd string, args ...string) (string, error) {
	var out []byte
	command := exec.Command(cmd, args...)
	out, err := command.CombinedOutput()

	return string(out), err
}
